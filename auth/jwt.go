package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jonDufty/recipes/auth/models/user"
)

const TOKEN_LIFETIME = time.Minute * 10

const JWT_KEY = "secret_key"

type UserInfoPayload struct {
	UserID   int64
	FullName string
}

type TokenClaims struct {
	*jwt.RegisteredClaims
	*UserInfoPayload
}

func GenerateToken(u *user.User) (string, error) {
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims = &TokenClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_LIFETIME)),
		},
		&UserInfoPayload{
			int64(u.ID),
			u.FullName,
		},
	}

	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(JWT_KEY))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(JWT_KEY), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return token, nil
}
