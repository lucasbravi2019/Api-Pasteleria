package mapper

import (
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
)

type RecipeMapper struct{}

type RecipeMapperInterface interface {
	RecipeDTOToRecipe(dto *dto.RecipeDTO) models.Recipe
}

var RecipeMapperInstance *RecipeMapper

func (m *RecipeMapper) RecipeDTOToRecipe(dto *dto.RecipeDTO) *models.Recipe {
	ingredients := []models.RecipeIngredient{}

	for i := 0; i < len(dto.Ingredients); i++ {
		ingredient := &models.RecipeIngredient{
			ID:       dto.Ingredients[i].ID,
			Name:     dto.Ingredients[i].Name,
			Package:  models.RecipeIngredientPackage(dto.Ingredients[i].Package),
			Quantity: float32(dto.Ingredients[i].Quantity),
			Price:    dto.Ingredients[i].Price,
		}

		ingredients = append(ingredients, *ingredient)
	}

	return &models.Recipe{
		ID:          dto.ID,
		Name:        dto.Name,
		Ingredients: ingredients,
		Price:       dto.Price,
	}
}
