package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
)

type RecipeIngredientMapper struct {
	RecipeMapper *RecipeMapper
}

var RecipeIngredientMapperInstance *RecipeIngredientMapper

func (m *RecipeIngredientMapper) ToRecipeIngredientDTO(id sql.NullInt64, name sql.NullString, quantity sql.NullFloat64,
	price sql.NullFloat64, pkg *dto.IngredientPackageDTO) *dto.RecipeIngredientDTO {
	if !id.Valid {
		return nil
	}

	return dto.NewRecipeIngredientDTO(db.GetLong(id), db.GetString(name), db.GetFloat(quantity), db.GetFloat(price), pkg)
}
