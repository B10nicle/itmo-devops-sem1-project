package server

import (
	"net/http"
	"project_sem/internal/database"
	"project_sem/internal/resource"
)

func NewServerRouter(repository *database.Repository) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v0/prices", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			resource.GetItem(repository)(w, r)
		case http.MethodPost:
			resource.CreateItem(repository)(w, r)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
	return mux
}
