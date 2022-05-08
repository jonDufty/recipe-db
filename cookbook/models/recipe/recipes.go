package recipe

import (
	"context"
	"errors"

	"github.com/jonDufty/recipes/common/database"
	"github.com/russross/meddler"
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

func GetAllRecipes(ctx context.Context) ([]*RecipeDB, error) {
	tx, err := database.StartTx(ctx)
	if err != nil {
		return nil, err
	}

	r := []*RecipeDB{}
	err = meddler.QueryAll(tx, r, "SELECT * from recipes")
	if err != nil {
		e := tx.Rollback()
		return nil, errors.New(e.Error() + e.Error())
	} else {
		err = tx.Commit()
	}

	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetRecipeById(ctx context.Context, id int) (*RecipeDB, error) {
	tx, err := database.StartTx(ctx)
	if err != nil {
		return nil, err
	}

	r := &RecipeDB{}
	err = meddler.QueryRow(tx, r, "SELECT * from recipes WHERE recipe_id = ?", id)
	if err != nil {
		e := tx.Rollback()
		return nil, errors.New(e.Error() + e.Error())
	} else {
		err = tx.Commit()
	}

	if err != nil {
		return nil, err
	}

	return r, nil
}

func GetRecipeByTitle(ctx context.Context, title string) (*RecipeDB, error) {
	tx, err := database.StartTx(ctx)
	if err != nil {
		return nil, err
	}

	r := &RecipeDB{}
	err = meddler.QueryRow(tx, r, "SELECT * from recipes WHERE title = ?", title)
	if err != nil {
		e := tx.Rollback()
		return nil, errors.New(e.Error() + e.Error())
	} else {
		err = tx.Commit()
	}

	if err != nil {
		return nil, err
	}

	return r, nil
}

func (r *RecipeDB) InsertRecipe(ctx context.Context) error {
	err := database.Insert(ctx, "recipes", r)
	return err
}

func (ri *RecipeIngredientDB) InsertRecipeIngredient(ctx context.Context) error {
	err := database.Insert(ctx, "recipe_instructions", ri)
	return err
}

func (ri *RecipeInstructionDB) InsertRecipeInstruction(ctx context.Context) error {
	err := database.Insert(ctx, "recipe_instructions", ri)
	return err
}
