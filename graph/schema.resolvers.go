package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"strconv"

	"github.com/jonDufty/recipes/cookbook/rpc/cookbookpb"
	"github.com/jonDufty/recipes/graph/generated"
	"github.com/jonDufty/recipes/graph/model"
)

func (r *ingredientResolver) Amount(ctx context.Context, obj *cookbookpb.Ingredient) (float64, error) {
	return float64(obj.Amount), nil
}

func (r *mutationResolver) CreateRecipe(ctx context.Context, input model.RecipeInput) (*cookbookpb.InsertRecipeResponse, error) {
	req := &cookbookpb.Recipe{
		Title:        input.Title,
		Description:  input.Description,
		Ingredients:  input.Ingredients,
		Instructions: input.Instructions,
	}
	resp, err := r.Clients.Cookbook.InsertRecipe(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *queryResolver) Recipes(ctx context.Context, id *string) ([]*cookbookpb.Recipe, error) {
	var recipes []*cookbookpb.Recipe
	if id != nil {
		reqId, err := strconv.Atoi(*id)
		if err != nil {
			return nil, errors.New("failed to convert id to integer")
		}
		resp, err := r.Clients.Cookbook.GetRecipeById(
			ctx,
			&cookbookpb.GetRecipeByIdRequest{
				Id: int64(reqId),
			})

		if err != nil {
			return nil, err
		}
		recipes = append(recipes, resp)

	} else {
		req := &cookbookpb.ListRecipesRequest{}
		resp, err := r.Clients.Cookbook.ListRecipes(ctx, req)
		if err != nil {
			return nil, err
		}

		recipes = resp.Recipes
	}
	return recipes, nil
}

func (r *queryResolver) RecipesAuth(ctx context.Context, id *string) ([]*cookbookpb.Recipe, error) {
	return r.Recipes(ctx, id)
}

func (r *ingredientInputResolver) Amount(ctx context.Context, obj *cookbookpb.Ingredient, data float64) error {
	if obj.Amount < 0 {
		return errors.New("Amount must be greater than 0")
	}
	return nil
}

// Ingredient returns generated.IngredientResolver implementation.
func (r *Resolver) Ingredient() generated.IngredientResolver { return &ingredientResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// IngredientInput returns generated.IngredientInputResolver implementation.
func (r *Resolver) IngredientInput() generated.IngredientInputResolver {
	return &ingredientInputResolver{r}
}

type ingredientResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type ingredientInputResolver struct{ *Resolver }
