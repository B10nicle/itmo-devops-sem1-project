package database

import (
	"database/sql"
	"fmt"
	"time"
)

type Item struct {
	ID         int
	Name       string
	Category   string
	Price      float64
	CreateDate time.Time
}

type FilterParams struct {
	MinPrice      float64
	MaxPrice      float64
	MinCreateDate time.Time
	MaxCreateDate time.Time
}

func (r *Repository) CreateItem(tx *sql.Tx, item Item) error {
	query := `
		INSERT INTO prices (name, category, price, create_date) 
		VALUES ($1, $2, $3, $4)`
	_, err := tx.Exec(query, item.Name, item.Category, item.Price, item.CreateDate)
	if err != nil {
		return fmt.Errorf("failed to create price: %w", err)
	}
	return nil
}

func (r *Repository) GetItems(params FilterParams) ([]Item, error) {
	query := `
		SELECT id, name, category, price, create_date 
		FROM prices 
		WHERE price >= $1 AND price <= $2 
		AND create_date BETWEEN $3 AND $4`
	rows, err := r.db.Query(query, params.MinPrice, params.MaxPrice, params.MinCreateDate, params.MaxCreateDate)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve prices: %w", err)
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	prices := make([]Item, 0)
	for rows.Next() {
		var price Item
		err = rows.Scan(&price.ID, &price.Name, &price.Category, &price.Price, &price.CreateDate)
		if err != nil {
			return nil, fmt.Errorf("failed to scan price row: %w", err)
		}
		prices = append(prices, price)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("row iteration error: %w", err)
	}

	return prices, nil
}

func (r *Repository) GetUniqueCategoriesAndTotalPrice() (float64, int, error) {
	var totalPrice float64
	var totalCategories int

	query := "SELECT SUM(price), COUNT(DISTINCT category) FROM prices"
	err := r.db.QueryRow(query).Scan(&totalPrice, &totalCategories)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to calculate totals: %w", err)
	}

	return totalPrice, totalCategories, nil
}
