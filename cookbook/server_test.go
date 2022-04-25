package cookbook

import (
	"net/http/httptest"
	"testing"

	"github.com/jonDufty/recipes/config"
	"github.com/stretchr/testify/require"
)

func newTestDB() *TestApp {
	cfg := config.NewCookbookConfig()
	testApp := NewTestApp(cfg)

	testApp.InitDB()

	return testApp
}

func TestListRecipes(t *testing.T) {
	testApp := newTestDB()
	testApp.InitServers()

	err := testApp.App.DB.Ping()
	require.NoError(t, err)

	testApp.PopulateTestCookbook()
	// ctx := testApp.App.Ctx

	testServer := httptest.NewServer(testApp.Http)

	defer testServer.Close()
	defer testApp.Closer()

	type testCase struct {
	}
}

func TestInsertRecipe(t *testing.T) {

}

func TestGetRecipe(t *testing.T) {

}
