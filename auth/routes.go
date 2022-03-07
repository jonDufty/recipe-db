package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	authpb "github.com/jonDufty/recipes/auth/rpc/authpb"

	"github.com/jonDufty/recipes/auth/models/session"
	"github.com/jonDufty/recipes/auth/models/user"

	"github.com/jonDufty/recipes/common/crypto"
	db "github.com/jonDufty/recipes/common/database"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type AuthPostParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	FullName string `json:"full_name,omitempty"`
}

// Returns chi router with common middleware applied
// Middleware includes:
// 		Logger
func NewRouter(a *App) http.Handler {
	r := chi.NewRouter()

	// Add common middleware
	r.Use(middleware.Logger)
	r.Use(db.Middleware(a.DB))

	r.Get("/system/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "🔐")
	})

	// Common auth endpoints
	r.Group(func(r chi.Router) {
		// Other middleware

		r.Post("/login", login)
		r.Post("/register", register)

	})

	// Test authenticated endpoints
	r.Group(func(r chi.Router) {
		r.Use(session.Middleware())

		r.Get("/check", check)
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
		return
	}

	hash, err := crypto.HashPassword(d.Password)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	requestUser := &user.User{
		FullName:     d.FullName,
		Email:        d.Email,
		TimeCreated:  time.Now(),
		PasswordHash: hash,
	}

	err = requestUser.InsertUser(r.Context())
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Printf("Error inserting user into db: %v", requestUser)
	}
	log.Print("User created successfully")

}

func login(w http.ResponseWriter, r *http.Request) {

	d, err := getPostParams(r)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Print(err.Error())
		return
	}

	u, err := user.GetUserByEmail(r.Context(), d.Email)
	if err != nil {
		http.Error(w, "Incorrect email or password", http.StatusBadRequest)
		log.Printf("User of email: %s not found", d.Email)
		log.Print(err.Error())
		return
	}

	err = crypto.CheckPassword(d.Password, u.PasswordHash)
	if err != nil {
		http.Error(w, "Incorrect email or password", http.StatusBadRequest)
		log.Printf("Password for email: %s does not match hash\nExpected: %s\n Received %s", d.Email, u.PasswordHash, d.Password)
		log.Print(err.Error())
		return
	}

	err = session.CreateSession(w, r, u)
	if err != nil {
		http.Error(w, "Error creating session", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/system/healthcheck", http.StatusOK)
}

func check(w http.ResponseWriter, r *http.Request) {
	err := session.CheckCookie(w, r)

	if err != nil {
		fmt.Fprintf(w, "You are not logged in...")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "You are logged in")
	w.WriteHeader(http.StatusOK)
}
