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
		log.Println("Warning: .env file not found, using environment variables.")
	}

	return Config{
		Server: LoadServerConfig(),
		DB:     LoadDBConfig(),
	}
}
