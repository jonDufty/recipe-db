package cookbook

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

func NewTestApp(c *config.CookbookConfig) *TestApp {

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
	ta.App.Ctx = database.DbContext(context.Background(), ta.App.DB)
}

func (ta *TestApp) InitServers() {
	ta.Twirp = NewServer(ta.App)
	ta.Http = NewRouter(ta.App)
}

func (ta *TestApp) PopulateTestCookbook() {

}
