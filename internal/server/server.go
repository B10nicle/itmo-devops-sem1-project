package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project_sem/internal/config"
	"project_sem/internal/database"
	"syscall"
	"time"
)

type App struct {
	server          *http.Server
	shutdownTimeout time.Duration
}

func New(cfg config.Config) (*App, error) {
	repo, err := database.NewRepository(cfg.DB)
	if err != nil {
		return nil, fmt.Errorf("failed to create repository: %w", err)
	}

	router := NewServerRouter(repo)
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Server.Port),
		Handler:      router,
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
	}

	return &App{
		server:          server,
		shutdownTimeout: 5 * time.Second,
	}, nil
}

func (app *App) Run() {
	go func() {
		log.Printf("Starting server on port %s", app.server.Addr)
		if err := app.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("Server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Received signal: %s. Initiating shutdown\n", sig)

	ctx, cancel := context.WithTimeout(context.Background(), app.shutdownTimeout)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		log.Printf("Error during server shutdown: %s", err)
	} else {
		log.Println("Server shutdown completed successfully.")
	}
}
