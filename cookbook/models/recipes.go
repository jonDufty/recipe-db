package models

import (
	"context"

	"github.com/jonDufty/recipes/common/database"
)

type Ingredient struct {
	Id    int    `json:"Id" meddler:"ingredient_id,pk"`
	Label string `json:"label" meddler:"label"`
}
type Unit struct {
	Id    int    `json:"id" meddler:"unit_id,pk"`
	Label string `json:"label" meddler:"label"`
}

type RecipeIngredientDB struct {
	Id           int     `json:"id" meddler:"recipe_ingredient_id,pk"`
	IngredientId int     `json:"ingredient" meddler:"ingredient_id"`
	UnitId       int     `json:"unit" meddler:"unit_id"`
	Amount       float32 `json:"amount" meddler:"amount"`
	RecipeId     int     `json:"recipe_id" meddler:"recipe_id"`
}

type RecipeInstructionDB struct {
	ID     int    `json:"id" meddler:"instruction_id,pk"`
	Step   int    `json:"step" meddler:"step"`
	Text   string `json:"text" meddler:"text"`
	Recipe int    `json:"recipe_id" meddler:"recipe_id"`
}

type RecipeDB struct {
	Id          int    `json:"id" meddler:"recipe_id,pk"`
	Title       string `json:"title" meddler:"title"`
	Description string `json:"description" meddler:"description"`
}

func NewMockIngredient() *Ingredient {
	return &Ingredient{
		Id:    1,
		Label: "Ingredient",
	}
}

func NewMockUnit() *Unit {
	return &Unit{
		Id:    1,
		Label: "Unit",
	}
}

func NewMockRecipeIngredient() *RecipeIngredientDB {
	return &RecipeIngredientDB{
		Id:           1,
		IngredientId: 1,
		UnitId:       1,
		Amount:       1.0,
		RecipeId:     1,
	}
}

func NewMockInstruction() *RecipeInstructionDB {
	return &RecipeInstructionDB{
		ID:     1,
		Step:   1,
		Text:   "text",
		Recipe: 1,
	}
}

func NewMockRecipe() *RecipeDB {
	return &RecipeDB{
		Id:          1,
		Title:       "Recipe",
		Description: "Description",
	}
}

func GetRecipeById(ctx context.Context, id int) *RecipeDB {
	panic("implement me")
}

func GetRecipeByTitle(ctx context.Context, title string) *RecipeDB {
	panic("implement me")
}

func GetIngredientById(ctx context.Context, id int) *Ingredient {
	panic("implement me")
}

func GetUnitById(ctx context.Context, id int) *Unit {
	panic("implement me")
}

func GetRecipeIngredientById(ctx context.Context, id int) *RecipeIngredientDB {
	panic("implement me")
}

func (r *RecipeDB) GetRecipeIngredients(ctx context.Context) []RecipeIngredientDB {
	panic("implement me")
}

func (r *RecipeDB) GetRecipeInstructions(ctx context.Context) []RecipeInstructionDB {
	panic("implement me")
}

func (ri *RecipeIngredientDB) InsertRecipeIngredient(ctx context.Context) error {
	err := database.Insert(ctx, "recipe_instructions", ri)
	return err
}

func (ri *RecipeInstructionDB) InsertRecipeInstruction(ctx context.Context) error {
	err := database.Insert(ctx, "recipe_instructions", ri)
	return err
}
