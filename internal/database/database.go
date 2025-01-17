package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"project_sem/internal/config"
	"time"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(cfg config.DB) (*Repository, error) {
	log.Println("Connecting to the database")

	connStr := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Printf("Successfully connected to database '%s'", cfg.Name)
	return &Repository{db: db}, nil
}

func (r *Repository) Close() {
	if err := r.db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	} else {
		log.Println("Database connection closed successfully.")
	}
}
