package auth

import (
	"context"
	"net/http"

	"github.com/jonDufty/recipes/common/database"
	"github.com/jonDufty/recipes/config"
)

type TestApp struct {
	App    *App
	Twirp  *Server
	Http   http.Handler
	Closer func()
}

// var users []*user.User = []*user.User{
// 	&user.User{
// 		ID:           1,
// 		FullName:     "Test User",
// 		Email:        "test1@example.com",
// 		PasswordHash: "password123",
// 	},
// 	&user.User{
// 		ID:           2,
// 		FullName:     "Test User",
// 		Email:        "test2@example.com",
// 		PasswordHash: "password123",
// 	},
// }

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
	db, closer := database.NewTestDBConnection()
	ta.App.DB = db
	ta.Closer = closer
}

func (ta *TestApp) InitServers() {
	ta.Twirp = NewServer(ta.App)
	ta.Http = NewRouter(ta.App)
}

// func (ta *TestApp) PopulateTestUsers() {
// 	for _, u := range users {
// 		u.InsertUser(ta.App.Ctx)
// 	}
// }
