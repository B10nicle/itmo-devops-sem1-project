package services

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"project_sem/internal/archiver"
	"project_sem/internal/database"
	"project_sem/internal/serializers"
)

type PriceStats struct {
	TotalCount      int `json:"total_count"`
	DuplicateCount  int `json:"duplicates_count"`
	TotalItems      int `json:"total_items"`
	TotalCategories int `json:"total_categories"`
	TotalPrice      int `json:"total_price"`
}

func CreatePrice(repo *database.Repository) http.HandlerFunc {
	const errorResponseBody = "failed to upload prices"
	const successContentType = "application/json"

	return func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			log.Printf("failed to read incoming file: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				log.Printf("failed to close file: %v\n", err)
			}
		}(file)

		formatType := r.URL.Query().Get("type")
		rc, err := extractFile(file, formatType)
		if err != nil {
			log.Printf("failed to unarchive incoming file: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		defer func(rc io.ReadCloser) {
			err := rc.Close()
			if err != nil {
				log.Printf("failed to close reader: %v\n", err)
			}
		}(rc)

		// Deserialize prices
		prices, deserializationErrors := serializers.DeserializePrices(rc)
		totalCount := len(prices) + len(deserializationErrors)

		if len(deserializationErrors) > 0 {
			log.Printf("deserialization errors encountered: %v\n", deserializationErrors)
		}

		stats := PriceStats{
			TotalCount: totalCount,
		}

		for _, price := range prices {
			err = repo.CreatePrice(price)
			if err != nil {
				stats.DuplicateCount++
			} else {
				stats.TotalItems++
			}
		}

		totalPrice, totalCategories, err := repo.GetUniqueCategoriesAndTotalPrice()
		if err != nil {
			log.Printf("failed to get total price and unique categories: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
		stats.TotalCategories = totalCategories
		stats.TotalPrice = int(totalPrice)

		w.Header().Set("Content-Type", successContentType)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(stats); err != nil {
			log.Printf("failed to encode response: %v\n", err)
			http.Error(w, "failed to encode response", http.StatusInternalServerError)
		}
	}
}

func extractFile(r io.Reader, fileType string) (io.ReadCloser, error) {
	switch fileType {
	case "zip":
		return archiver.ExtractFromZip(r)
	case "tar":
		return archiver.ExtractFromTar(r)
	default:
		return nil, errors.New("unsupported archiver type")
	}
}
