package cookbook

import (
	"context"
	"net/http"

	"github.com/jonDufty/recipes/common/database"
	"github.com/jonDufty/recipes/config"
	"github.com/jonDufty/recipes/cookbook/models/recipe"
)

type TestApp struct {
	App    *App
	Twirp  *Server
	Http   http.Handler
	Closer func()
}

var testRecipes []*recipe.RecipeDB = []*recipe.RecipeDB{
	{
		Title:       "Test recipe 1",
		Description: "A chicken curry",
	},
	{
		Title:       "Test recipe 2",
		Description: "A lamb curry",
	},
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

func (ta *TestApp) InitDB(dbName string) {
	db, closer := database.NewTestDBConnection(dbName)
	ta.App.DB = db
	ta.Closer = closer
	ta.App.Ctx = database.DbContext(context.Background(), ta.App.DB)
}

func (ta *TestApp) InitServers() {
	ta.Twirp = NewServer(ta.App)
	ta.Http = NewRouter(ta.App)
}

func (ta *TestApp) PopulateTestCookbook() {

	for _, r := range testRecipes {
		r.InsertRecipe(ta.App.Ctx)
	}
}
