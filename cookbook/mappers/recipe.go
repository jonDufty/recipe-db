package mappers

import (
	"github.com/jonDufty/recipes/cookbook/models"
)

type RecipeIngredientAPI struct {
	Ingredient models.Ingredient `json:"ingredient"`
	Unit       models.Unit       `json:"unit"`
	Amount     float32           `json:"amount"`
}

type RecipeInstructionAPI struct {
	Step int    `json:"step"`
	Text string `json:"text"`
}

type RecipeAPI struct {
	Id           int                    `json:"id"`
	Title        string                 `json:"title"`
	Description  string                 `json:"description"`
	Ingredients  []RecipeIngredientAPI  `json:"ingredients"`
	Instructions []RecipeInstructionAPI `json:"instructions"`
}

func RecipeFromDbModel(model *models.RecipeDB) *RecipeAPI {
	return &RecipeAPI{
		Id:           model.Id,
		Title:        model.Title,
		Description:  model.Description,
		Ingredients:  []RecipeIngredientAPI{},
		Instructions: []RecipeInstructionAPI{},
	}
}
