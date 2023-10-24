package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

func ToRecipeIngredientList(rows *sql.Rows) (*[]models.RecipeIngredient, error) {
	recipes := util.NewList[models.RecipeIngredient]()

	for rows.Next() {
		var recipeId int64
		var recipeName string
		var recipePrice float64
		var ingredientQuantityUsed sql.NullFloat64
		var ingredientId sql.NullInt64
		var ingredientUsedPrice sql.NullFloat64
		var ingredientPackagePrice sql.NullFloat64
		var packageId sql.NullInt64
		var metric sql.NullString
		var packageQuantity sql.NullFloat64

		err := rows.Scan(&recipeId, &recipeName, &recipePrice, &ingredientQuantityUsed, &ingredientId, &ingredientUsedPrice,
			&ingredientPackagePrice, &packageId, &metric, &packageQuantity)

		if pkg.HasError(err) {
			return nil, err
		}

		var ingredientPackage *models.RecipeIngredientPackage
		var recipe *models.RecipeIngredient

		if packageId.Valid {
			ingredientPackage = models.NewRecipeIngredientPackage(db.GetLong(packageId), db.GetString(metric), db.GetFloat(packageQuantity),
				db.GetFloat(ingredientPackagePrice))
		}

		if ingredientId.Valid {
			recipe = models.NewRecipeIngredient(recipeId, recipeName, ingredientPackage, ingredientQuantityUsed.Float64,
				db.GetFloat(ingredientUsedPrice))
		}

		if recipe != nil {
			util.Add(&recipes, *recipe)
		}
	}

	return &recipes, nil
}

func ToRecipeIngredientDTOList(ingredients *[]models.RecipeIngredient) *[]dto.RecipeIngredientDTO {
	recipeIngredients := util.NewList[dto.RecipeIngredientDTO]()

	for _, ingredient := range *ingredients {
		ingredientPackage := dto.NewPackageDTO(ingredient.Package.PackageId, ingredient.Package.Metric,
			ingredient.Package.Quantity, *ingredient.Price)
		recipeIngredient := dto.NewRecipeIngredientDTO(ingredient.IngredientId, ingredient.Name, ingredient.Quantity,
			*ingredientPackage)

		util.Add(&recipeIngredients, *recipeIngredient)
	}

	return &recipeIngredients
}
