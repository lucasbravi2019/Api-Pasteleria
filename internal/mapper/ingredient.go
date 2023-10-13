package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
)

func ToIngredientList(rows *sql.Rows) []models.Ingredient {

	return nil
}

func ToIngredient(rows *sql.Rows) *models.Ingredient {

	return nil
}

func ToIngredientDTO(rows *sql.Rows) *dto.IngredientDTO {
	return nil
}

func ToIngredientDTOList(rows *sql.Rows) *[]dto.IngredientDTO {
	return nil
}
