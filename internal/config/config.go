package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Server Server
	DB     DB
}

func Load() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return Config{
		Server: LoadServerConfig(),
		DB:     LoadDBConfig(),
	}
}
