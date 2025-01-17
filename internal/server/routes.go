package server

import (
	"net/http"
	"project_sem/internal/database"
	"project_sem/internal/services"
)

func NewServerRouter(repo *database.Repository) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v0/prices", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			services.GetPrice(repo)(w, r)
		case http.MethodPost:
			services.CreatePrice(repo)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}
