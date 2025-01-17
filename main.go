package main

import (
	"log"
	"project_sem/internal/config"
	"project_sem/internal/server"
)

func main() {
	log.Println("Start loading server and database configuration")
	cfg := config.Load()
	log.Println("Server and database configuration has been successfully loaded")

	application, err := server.New(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize the application: %v", err)
	}

	application.Run()
}
