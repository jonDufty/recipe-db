package cookbook

import (
	"context"

	"github.com/jonDufty/recipes/cookbook/mappers"
	"github.com/jonDufty/recipes/cookbook/models/recipe"
	rpc "github.com/jonDufty/recipes/cookbook/rpc/cookbookpb"
	"github.com/twitchtv/twirp"
)

type Server struct {
	app *App
}

func NewServer(a *App) *Server {
	return &Server{a}
}

func (s *Server) GetRecipeById(ctx context.Context, r *rpc.GetRecipeByIdRequest) (*rpc.Recipe, error) {
	recipe, err := recipe.GetRecipeById(ctx, int(r.Id))
	if err != nil {
		return nil, err
	}
	return mappers.PbRecipeFromDbModel(recipe), nil
}

func (s *Server) InsertRecipe(ctx context.Context, r *rpc.Recipe) (*rpc.InsertRecipeResponse, error) {
	model := mappers.DbModelFromRpc(r)
	err := model.InsertRecipe(ctx)
	if err != nil {
		return nil, err
	}

	return &rpc.InsertRecipeResponse{
		Id:     int64(model.Id),
		Errors: string(twirp.NoError),
	}, nil
}

func (s *Server) ListRecipes(ctx context.Context, r *rpc.ListRecipesRequest) (*rpc.RecipeList, error) {
	recipes, err := recipe.GetAllRecipes(ctx)
	if err != nil {
		return nil, err
	}

	apiRecipes := make([]*rpc.Recipe, len(recipes))

	for idx, r := range recipes {
		apiRecipes[idx] = mappers.PbRecipeFromDbModel(r)
	}

	return &rpc.RecipeList{
		Recipes: apiRecipes,
	}, nil
}
