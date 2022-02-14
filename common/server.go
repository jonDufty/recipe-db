package common

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router struct{}

// Returns chi router with common middleware applied
// Middleware includes:
// 		Logger
func NewRouter() *chi.Mux {
	r := chi.NewRouter()

	// Add common middleware
	r.Use(middleware.Logger)

	return r
}
