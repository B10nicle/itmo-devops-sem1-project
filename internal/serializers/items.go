package serializers

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"project_sem/internal/database"
	"strconv"
	"time"
)

func SerializeItems(items []database.Item) (*bytes.Buffer, error) {
	var buffer bytes.Buffer
	csvWriter := csv.NewWriter(&buffer)
	defer csvWriter.Flush()

	if err := csvWriter.Write([]string{"id", "name", "category", "price", "create_date"}); err != nil {
		return nil, fmt.Errorf("failed to write header: %w", err)
	}

	for _, price := range items {
		record := []string{
			fmt.Sprintf("%d", price.ID),
			price.Name,
			price.Category,
			fmt.Sprintf("%.2f", price.Price),
			price.CreateDate.Format("2006-01-02"),
		}
		if err := csvWriter.Write(record); err != nil {
			return nil, fmt.Errorf("failed to write record for price ID %d: %w", price.ID, err)
		}
	}

	return &buffer, nil
}

func DeserializeItems(r io.Reader) ([]database.Item, []error) {
	var items []database.Item
	var errors []error

	csvReader := csv.NewReader(r)

	headerRead := false
	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to read record: %w", err))
			continue
		}

		if !headerRead {
			headerRead = true
			continue
		}

		item, err := validateItem(record)
		if err != nil {
			errors = append(errors, fmt.Errorf("invalid record %v: %w", record, err))
			continue
		}

		items = append(items, item)
	}

	return items, errors
}

func validateItem(record []string) (database.Item, error) {
	if len(record) != 5 {
		return database.Item{}, errors.New("invalid number of fields in record")
	}

	id, err := strconv.Atoi(record[0])
	if err != nil {
		return database.Item{}, fmt.Errorf("invalid ID: %w", err)
	}

	name := record[1]
	if name == "" {
		return database.Item{}, errors.New("name cannot be empty")
	}

	category := record[2]
	if category == "" {
		return database.Item{}, errors.New("category cannot be empty")
	}

	price, err := strconv.ParseFloat(record[3], 64)
	if err != nil {
		return database.Item{}, fmt.Errorf("invalid price: %w", err)
	}

	createDate, err := time.Parse("2006-01-02", record[4])
	if err != nil {
		return database.Item{}, fmt.Errorf("invalid create date: %w", err)
	}

	return database.Item{
		ID:         id,
		Name:       name,
		Category:   category,
		Price:      price,
		CreateDate: createDate,
	}, nil
}
