package config

import (
	"os"
	"project_sem/internal/utils"
)

type DB struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
}

func LoadDBConfig() DB {
	return DB{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     utils.ParseInt(os.Getenv("POSTGRES_PORT")),
		Name:     os.Getenv("POSTGRES_DB"),
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}
}
