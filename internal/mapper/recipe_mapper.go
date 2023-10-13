package mapper

import (
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
)

type RecipeMapper struct{}

type RecipeMapperInterface interface {
	RecipeDTOToRecipe(dto *dto.RecipeDTO) models.Recipe
}

var RecipeMapperInstance *RecipeMapper

func (m *RecipeMapper) RecipeDTOToRecipe(dto *dto.RecipeDTO) *models.Recipe {

	return nil
}
