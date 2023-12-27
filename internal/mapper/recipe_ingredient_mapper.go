package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type RecipeIngredientMapper struct {
	RecipeMapper *RecipeMapper
}

var RecipeIngredientMapperInstance *RecipeIngredientMapper

type RecipeIngredientMapperInterface interface {
	ToRecipeIngredientList(rows *sql.Rows) (*[]models.RecipeIngredient, error)

	ToRecipeIngredientDTOList(ingredients *[]models.RecipeIngredient) *[]dto.RecipeIngredientDTO

	ToRecipeIngredient(ingredientId sql.NullInt64, ingredientName sql.NullString,
		ingredientQuantity sql.NullFloat64, ingredientPrice sql.NullFloat64,
		recipeId sql.NullInt64) *models.RecipeIngredient

	ToRecipeIngredientDTO(id int64, name string, quantity float64, pkg dto.PackageDTO) *dto.RecipeIngredientDTO

	SetPackageToRecipeIngredientDTO(recipeIngredient *models.RecipeIngredient, ingredientPackage *models.RecipeIngredientPackage) *models.RecipeIngredient
}

func (m *RecipeIngredientMapper) ToRecipeIngredientList(rows *sql.Rows) (*[]models.RecipeIngredient, error) {
	recipes := util.NewList[models.RecipeIngredient]()

	for rows.Next() {
		var recipeId int64
		var ingredientQuantityUsed sql.NullFloat64
		var ingredientId sql.NullInt64
		var ingredientUsedPrice sql.NullFloat64
		var ingredientPackagePrice sql.NullFloat64
		var packageId sql.NullInt64
		var metric sql.NullString
		var packageQuantity sql.NullFloat64

		err := rows.Scan(&recipeId, &ingredientQuantityUsed, &ingredientId, &ingredientUsedPrice,
			&ingredientPackagePrice, &packageId, &metric, &packageQuantity)

		if pkg.HasError(err) {
			return nil, err
		}

	}

	return &recipes, nil
}

func (m *RecipeIngredientMapper) ToRecipeIngredientDTOList(ingredients *[]models.RecipeIngredient) *[]dto.RecipeIngredientDTO {
	recipeIngredients := util.NewList[dto.RecipeIngredientDTO]()

	for _, ingredient := range *ingredients {
		ingredientPackage := dto.NewPackageDTO(ingredient.Package.PackageId, ingredient.Package.Metric,
			ingredient.Package.Quantity, *ingredient.Price)
		recipeIngredient := m.ToRecipeIngredientDTO(*ingredient.IngredientId, *ingredient.Name, *ingredient.Quantity, *ingredientPackage)

		util.Add(&recipeIngredients, *recipeIngredient)
	}

	return &recipeIngredients
}

func (m *RecipeIngredientMapper) ToRecipeIngredient(ingredientId sql.NullInt64, ingredientName sql.NullString,
	ingredientQuantity sql.NullFloat64, ingredientPrice sql.NullFloat64,
	recipeId sql.NullInt64) *models.RecipeIngredient {
	return models.NewRecipeIngredient(
		db.GetLong(ingredientId),
		db.GetString(ingredientName),
		db.GetFloat(ingredientQuantity),
		db.GetFloat(ingredientPrice),
		db.GetLong(recipeId))
}

func (m *RecipeIngredientMapper) ToRecipeIngredientDTO(id int64, name string, quantity float64, pkg dto.PackageDTO) *dto.RecipeIngredientDTO {
	return dto.NewRecipeIngredientDTO(id, name, quantity, pkg)
}

func (m *RecipeIngredientMapper) SetPackageToRecipeIngredientDTO(recipeIngredient *models.RecipeIngredient,
	ingredientPackage *models.RecipeIngredientPackage) {
	recipeIngredient.Package = ingredientPackage
}
