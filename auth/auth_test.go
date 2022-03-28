package auth

import (
	"context"

	"github.com/jonDufty/recipes/auth/models/user"
	"github.com/jonDufty/recipes/common/database"
	"github.com/jonDufty/recipes/config"
)

var users []*user.User = []*user.User{
	&user.User{
		ID:           1,
		FullName:     "Test User",
		Email:        "test1@example.com",
		PasswordHash: "password123",
	},
	&user.User{
		ID:           2,
		FullName:     "Test User",
		Email:        "test2@example.com",
		PasswordHash: "password123",
	},
}

func NewTestApp(c *config.AuthConfig) *App {

	app := App{
		Config: c,
		Ctx:    context.Background(),
	}

	return &app
}

func AddTestAppDB(testApp *App) func() {
	db, closer := database.NewTestDBConnection()
	testApp.DB = db

	return closer
}

func PopulateTestUsers(testApp *App) {
	for _, u := range users {
		u.InsertUser(testApp.Ctx)
	}
}
