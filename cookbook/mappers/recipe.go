package mappers

import (
	"github.com/jonDufty/recipes/cookbook/models/recipe"
	rpc "github.com/jonDufty/recipes/cookbook/rpc/cookbookpb"
)

func PbRecipeFromDbModel(model *recipe.RecipeDB) *rpc.Recipe {
	return &rpc.Recipe{
		Id:           int64(model.Id),
		Title:        model.Title,
		Description:  model.Description,
		Ingredients:  []*rpc.Ingredient{},
		Instructions: []*rpc.Instruction{},
	}
}

func DbModelFromRpc(r *rpc.Recipe) *recipe.RecipeDB {
	return &recipe.RecipeDB{
		Id:          int(r.Id),
		Title:       r.Title,
		Description: r.Description,
	}
}
