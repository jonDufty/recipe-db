package auth

import (
	"fmt"
	"net/http"

	authpb "github.com/jonDufty/recipes/auth/rpc/authpb"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Returns chi router with common middleware applied
// Middleware includes:
// 		Logger
func NewRouter(a *App) http.Handler {
	r := chi.NewRouter()

	// Add common middleware
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🔐")
	})

	// Add twirp routes
	r.Group((func(chi.Router) {
		// Add middleware
		twirpServer := authpb.NewAuthServer(NewServer(a))
		r.Mount("/auth/twirp/", http.StripPrefix("/auth", twirpServer))
	}))

	return r
}
