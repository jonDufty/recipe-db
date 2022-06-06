package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/jonDufty/recipes/cookbook/rpc/cookbookpb"
	"github.com/jonDufty/recipes/graph/generated"
	"github.com/jonDufty/recipes/graph/model"
)

func (r *ingredientResolver) Amount(ctx context.Context, obj *cookbookpb.Ingredient) (float64, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := &model.Todo{
		Text: input.Text,
		ID:   fmt.Sprintf("T%d", rand.Int()),
		User: &model.User{ID: input.UserID, Name: "user " + input.UserID},
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) CreateRecipe(ctx context.Context, input cookbookpb.Recipe) (*cookbookpb.InsertRecipeResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) Recipes(ctx context.Context, id *string) ([]*cookbookpb.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) RecipesAuth(ctx context.Context, id *string) ([]*cookbookpb.Recipe, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *recipeResolver) Name(ctx context.Context, obj *cookbookpb.Recipe) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *ingredientInputResolver) Amount(ctx context.Context, obj *cookbookpb.Ingredient, data float64) error {
	panic(fmt.Errorf("not implemented"))
}

func (r *recipeInputResolver) Name(ctx context.Context, obj *cookbookpb.Recipe, data string) error {
	panic(fmt.Errorf("not implemented"))
}

// Ingredient returns generated.IngredientResolver implementation.
func (r *Resolver) Ingredient() generated.IngredientResolver { return &ingredientResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Recipe returns generated.RecipeResolver implementation.
func (r *Resolver) Recipe() generated.RecipeResolver { return &recipeResolver{r} }

// IngredientInput returns generated.IngredientInputResolver implementation.
func (r *Resolver) IngredientInput() generated.IngredientInputResolver {
	return &ingredientInputResolver{r}
}

// RecipeInput returns generated.RecipeInputResolver implementation.
func (r *Resolver) RecipeInput() generated.RecipeInputResolver { return &recipeInputResolver{r} }

type ingredientResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type recipeResolver struct{ *Resolver }
type ingredientInputResolver struct{ *Resolver }
type recipeInputResolver struct{ *Resolver }
