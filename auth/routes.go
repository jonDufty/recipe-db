package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	authpb "github.com/jonDufty/recipes/auth/rpc/authpb"

	"github.com/jonDufty/recipes/auth/models"
	"github.com/jonDufty/recipes/common/crypto"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type AuthPostParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
}

// Returns chi router with common middleware applied
// Middleware includes:
// 		Logger
func NewRouter(a *App) http.Handler {
	r := chi.NewRouter()

	// Add common middleware
	r.Use(middleware.Logger)

	r.Get("/system/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🔐")
	})

	// Common auth endpoints
	r.Group(func(r chi.Router) {
		// Other middleware

		r.Post("/login", login)
		r.Post("/register", register)

	})

	// Add twirp routes
	r.Group((func(chi.Router) {
		// Add middleware
		twirpServer := authpb.NewAuthServer(NewServer(a))
		r.Mount("/auth/twirp/", http.StripPrefix("/auth", twirpServer))
	}))

	return r
}

func getPostParams(r *http.Request) (AuthPostParams, error) {
	d := AuthPostParams{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&d)

	if err != nil {
		return d, err
	}

	return d, nil
}

func register(w http.ResponseWriter, r *http.Request) {

	d, err := getPostParams(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	hash, err := crypto.HashPassword(d.Password)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	requestUser := &models.User{
		FullName:     d.FullName,
		Email:        d.Email,
		TimeCreated:  time.Now(),
		PasswordHash: hash,
	}

	fmt.Print(d)
	fmt.Printf("%v", *requestUser)
	w.Write([]byte("Hello there"))

}

func login(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello there"))
}
