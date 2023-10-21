package mapper

import (
	"database/sql"

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
		var ingredientQuantityUsed float64
		var ingredientId int64
		var ingredientUsedPrice float64
		var ingredientPackagePrice float64
		var packageId int64
		var metric string
		var packageQuantity float64

		err := rows.Scan(&recipeId, &recipeName, &recipePrice, &ingredientQuantityUsed, &ingredientId, &ingredientUsedPrice,
			&ingredientPackagePrice, &packageId, &metric, &packageQuantity)

		if pkg.HasError(err) {
			return nil, err
		}

		ingredientPackage := models.NewRecipeIngredientPackage(ingredientId, metric, packageQuantity, ingredientPackagePrice)
		recipe := models.NewRecipeIngredient(recipeId, recipeName, *ingredientPackage, ingredientQuantityUsed, ingredientUsedPrice)

		util.Add(&recipes, *recipe)
	}

	return &recipes, nil
}

func ToRecipeIngredientDTOList(ingredients *[]models.RecipeIngredient) *[]dto.RecipeIngredientDTO {
	recipeIngredients := util.NewList[dto.RecipeIngredientDTO]()

	for _, ingredient := range *ingredients {
		ingredientPackage := dto.NewPackageDTO(ingredient.IngredientPackage.PackageId, ingredient.IngredientPackage.Metric,
			ingredient.IngredientPackage.Quantity, ingredient.Price)
		recipeIngredient := dto.NewRecipeIngredientDTO(ingredient.IngredientId, ingredient.IngredientName, ingredient.Quantity,
			*ingredientPackage)

		util.Add(&recipeIngredients, *recipeIngredient)
	}

	return &recipeIngredients
}
