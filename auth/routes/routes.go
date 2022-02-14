package routes

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Returns chi router with common middleware applied
// Middleware includes:
// 		Logger
func NewRouter() http.Handler {
	r := chi.NewRouter()

	// Add common middleware
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🔐")
	})

	return r
}
