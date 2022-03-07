package crypto

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(pwd string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(pwd), 14)

	if err != nil {
		return "", errors.New("Error hashing password." + err.Error())
	}

	return string(bytes), nil
}

func CheckPassword(pwd string, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pwd))
	return err
}
