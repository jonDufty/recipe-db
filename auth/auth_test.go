package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/jonDufty/recipes/auth/models/user"
	"github.com/jonDufty/recipes/common/crypto"
	"github.com/jonDufty/recipes/common/database"
	"github.com/jonDufty/recipes/config"
)

type TestApp struct {
	App    *App
	Twirp  *Server
	Http   http.Handler
	Closer func()
}

var users []*user.User = []*user.User{
	{
		FullName:     "Test User",
		Email:        "test1@example.com",
		PasswordHash: "password123",
		TimeCreated:  time.Now(),
	},
	{
		FullName:     "Test User",
		Email:        "test2@example.com",
		PasswordHash: "password123",
		TimeCreated:  time.Now(),
	},
}

func NewTestApp(c *config.AuthConfig) *TestApp {

	app := App{
		Config: c,
		Ctx:    context.Background(),
	}

	return &TestApp{
		App: &app,
	}
}

func (ta *TestApp) InitDB() {
	db, closer := database.NewTestDBConnection("test_recipes")
	ta.App.DB = db
	ta.Closer = closer
	ta.App.Ctx = database.DbContext(context.Background(), ta.App.DB)
}

func (ta *TestApp) InitServers() {
	ta.Twirp = NewServer(ta.App)
	ta.Http = NewRouter(ta.App)
}

func (ta *TestApp) PopulateTestUsers() {
	for _, u := range users {
		hash, _ := crypto.HashPassword(u.PasswordHash)
		u.PasswordHash = hash
		u.InsertUser(ta.App.Ctx)
	}
}
