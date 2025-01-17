package config

import (
	"os"
	"project_sem/internal/utils"
	"time"
)

type Server struct {
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func LoadServerConfig() Server {
	return Server{
		Port:         utils.ParseInt(os.Getenv("SERVER_PORT")),
		ReadTimeout:  utils.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT")),
		WriteTimeout: utils.ParseDuration(os.Getenv("SERVER_WRITE_TIMEOUT")),
	}
}
