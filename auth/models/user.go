package auth

import "time"

type User struct {
	ID           int       `json:"uid" meddler:"id,pk"`
	FullName     string    `json:"name" meddler:"name"`
	Email        string    `json:"email" meddler:"email"`
	TimeCreated  time.Time `json:"time_created" meddler:"time_created"`
	PasswordHash string    `json:"password_hash" meddler:"password_hash"`
}
