package cookbook

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jonDufty/recipes/config"
	rpc "github.com/jonDufty/recipes/cookbook/rpc/cookbookpb"
	"github.com/stretchr/testify/require"
)

func newTestDB(dbName string) *TestApp {
	cfg := config.NewCookbookConfig()
	testApp := NewTestApp(cfg)

	testApp.InitDB(dbName)
	testApp.PopulateTestCookbook()

	return testApp
}

func TestGetRecipe(t *testing.T) {
	testApp := newTestDB("test_get_recipes")
	testApp.InitServers()

	err := testApp.App.DB.Ping()
	require.NoError(t, err)

	testServer := httptest.NewServer(testApp.Http)

	defer testServer.Close()
	defer testApp.Closer()

	client := rpc.NewCookbookProtobufClient(testServer.URL+"/cookbook", &http.Client{})

	type testCase struct {
		name        string
		input       *rpc.GetRecipeByIdRequest
		expected    *rpc.Recipe
		shouldError bool
	}

	tests := []testCase{
		{
			"Test get recipe 1",
			&rpc.GetRecipeByIdRequest{
				Id: 1,
			},
			&rpc.Recipe{
				Id:          1,
				Title:       "Test recipe 1",
				Description: "A chicken curry",
			},
			false,
		},
		{
			"Test get recipe 2",
			&rpc.GetRecipeByIdRequest{
				Id: 2,
			},
			&rpc.Recipe{
				Id:          2,
				Title:       "Test recipe 2",
				Description: "A lamb curry",
			},
			false,
		},
		{
			"Test incorrect id",
			&rpc.GetRecipeByIdRequest{
				Id: 3,
			},
			nil,
			true,
		},
		{
			"Test invalid id",
			&rpc.GetRecipeByIdRequest{
				Id: -1,
			},
			nil,
			true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tc *testing.T) {
			resp, err := client.GetRecipeById(testApp.App.Ctx, test.input)
			if test.shouldError {
				require.Error(tc, err)
				return
			}
			require.NoError(tc, err)
			require.Equal(tc, test.expected.Description, resp.Description)
			require.Equal(tc, test.expected.Instructions, resp.Instructions)
		})
	}

	type listTestCase struct {
		name        string
		input       *rpc.ListRecipesRequest
		expected    *rpc.RecipeList
		ShouldError bool
	}

	recipes := []*rpc.Recipe{
		{
			Id:          1,
			Title:       "Test recipe 1",
			Description: "A chicken curry",
		},
		{
			Id:          2,
			Title:       "Test recipe 2",
			Description: "A lamb curry",
		},
	}

	listTestCases := []listTestCase{
		{
			"test list recipes",
			&rpc.ListRecipesRequest{},
			&rpc.RecipeList{
				Recipes: recipes,
			},
			false,
		},
	}

	for _, test := range listTestCases {
		t.Run(test.name, func(tx *testing.T) {
			resp, err := client.ListRecipes(testApp.App.Ctx, test.input)
			require.NoError(t, err)

			for idx, rec := range resp.Recipes {
				require.Equal(t, test.expected.Recipes[idx].Title, rec.Title)
				require.Equal(t, test.expected.Recipes[idx].Description, rec.Description)
			}
		})
	}
}

// func TestListRecipes(t *testing.T) {
// 	testApp := newTestDB("test_list_recipes")
// 	testApp.InitServers()

// 	err := testApp.App.DB.Ping()
// 	require.NoError(t, err)

// 	testServer := httptest.NewServer(testApp.Http)

// 	defer testServer.Close()
// 	defer testApp.Closer()

// 	client := rpc.NewCookbookProtobufClient(testServer.URL+"/cookbook", &http.Client{})

// 	type testCase struct {
// 		name        string
// 		input       *rpc.ListRecipesRequest
// 		expected    *rpc.RecipeList
// 		ShouldError bool
// 	}

// 	recipes := []*rpc.Recipe{
// 		{
// 			Id:          1,
// 			Title:       "Test recipe 1",
// 			Description: "A chicken curry",
// 		},
// 		{
// 			Id:          2,
// 			Title:       "Test recipe 2",
// 			Description: "A lamb curry",
// 		},
// 	}

// 	testCases := []testCase{
// 		{
// 			"test list recipes",
// 			&rpc.ListRecipesRequest{},
// 			&rpc.RecipeList{
// 				Recipes: recipes,
// 			},
// 			false,
// 		},
// 	}

// 	for _, test := range testCases {
// 		t.Run(test.name, func(tx *testing.T) {
// 			resp, err := client.ListRecipes(testApp.App.Ctx, test.input)
// 			require.NoError(t, err)

// 			for idx, rec := range resp.Recipes {
// 				require.Equal(t, test.expected.Recipes[idx].Title, rec.Title)
// 				require.Equal(t, test.expected.Recipes[idx].Description, rec.Description)
// 			}
// 		})
// 	}

// }
