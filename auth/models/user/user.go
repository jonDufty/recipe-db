package user

import (
	"context"
	"errors"
	"time"

	"github.com/jonDufty/recipes/common/database"
	"github.com/russross/meddler"
)

type User struct {
	ID           int       `json:"uid" meddler:"id,pk"`
	FullName     string    `json:"name" meddler:"name"`
	Email        string    `json:"email" meddler:"email"`
	TimeCreated  time.Time `json:"time_created" meddler:"time_created,utctimez"`
	PasswordHash string    `json:"password_hash" meddler:"password_hash"`
}

func NewTest() *User {
	return &User{
		ID:           123,
		FullName:     "Test user",
		Email:        "test@example.com",
		TimeCreated:  time.Now(),
		PasswordHash: "jsdnabfljvbnasflv",
	}
}

func (u *User) InsertUser(ctx context.Context) error {
	err := database.Insert(ctx, "user", u)
	return err
}

func GetUserByEmail(ctx context.Context, email string) (*User, error) {
	tx, err := database.StartTx(ctx)
	if err != nil {
		return nil, err
	}

	u := &User{}
	err = meddler.QueryRow(tx, u, "SELECT * from user WHERE email = ?", email)
	if err != nil {
		e := tx.Rollback()
		return nil, errors.New(e.Error() + e.Error())
	} else {
		err = tx.Commit()
	}

	if err != nil {
		return nil, err
	}

	return u, nil
}
