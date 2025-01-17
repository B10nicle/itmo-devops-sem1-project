package resource

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"project_sem/internal/archiver"
	"project_sem/internal/database"
	"project_sem/internal/serializers"
	"strconv"
	"time"
)

func GetItem(repo *database.Repository) http.HandlerFunc {
	const errorResponseBody = "Failed to load prices"
	const successContentType = "application/zip"
	const successContentDisposition = "attachment; filename=data.zip"
	const csvFileName = "data.csv"

	return func(w http.ResponseWriter, r *http.Request) {
		params, err := buildFilterParams(r)
		if err != nil {
			log.Printf("Invalid filter parameters: %v\n", err)
			http.Error(w, "Invalid filter parameters", http.StatusBadRequest)
			return
		}

		prices, err := repo.GetItems(params)
		if err != nil {
			log.Printf("Failed to fetch prices: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}

		serializedPrices, err := serializers.SerializeItems(prices)
		if err != nil {
			log.Printf("Failed to serialize prices: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", successContentType)
		w.Header().Set("Content-Disposition", successContentDisposition)

		if err := archiver.ZipFile(serializedPrices, w, csvFileName); err != nil {
			log.Printf("Failed to archiver prices: %v\n", err)
			http.Error(w, errorResponseBody, http.StatusInternalServerError)
			return
		}
	}
}

func buildFilterParams(r *http.Request) (database.FilterParams, error) {
	params := database.FilterParams{}

	startDateStr := r.URL.Query().Get("start")
	if startDateStr != "" {
		startDate, err := time.Parse("2006-01-02", startDateStr)
		if err != nil {
			return params, wrapError(err, "invalid start date format")
		}
		params.MinCreateDate = startDate
	} else {
		params.MinCreateDate = time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC) // Default to epoch start
	}

	endDateStr := r.URL.Query().Get("end")
	if endDateStr != "" {
		endDate, err := time.Parse("2006-01-02", endDateStr)
		if err != nil {
			return params, wrapError(err, "invalid end date format")
		}
		params.MaxCreateDate = endDate
	} else {
		params.MaxCreateDate = time.Now()
	}

	minPriceStr := r.URL.Query().Get("min")
	if minPriceStr != "" {
		minPrice, err := strconv.ParseFloat(minPriceStr, 64)
		if err != nil {
			return params, wrapError(err, "invalid minimum price")
		}
		params.MinPrice = minPrice
	}

	maxPriceStr := r.URL.Query().Get("max")
	if maxPriceStr != "" {
		maxPrice, err := strconv.ParseFloat(maxPriceStr, 64)
		if err != nil {
			return params, wrapError(err, "invalid maximum price")
		}
		params.MaxPrice = maxPrice
	} else {
		params.MaxPrice = math.MaxFloat64
	}

	return params, nil
}

func wrapError(err error, message string) error {
	return fmt.Errorf("%s: %w", message, err)
}
