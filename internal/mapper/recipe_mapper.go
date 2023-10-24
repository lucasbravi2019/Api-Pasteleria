package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

func ToRecipeList(rows *sql.Rows) *[]models.Recipe {
	recipesGrouper := util.NewMap[int64, models.Recipe]()

	for rows.Next() {
		var recipeId int64
		var recipeName string
		var recipePrice float64
		var ingredientId sql.NullInt64
		var ingredientName sql.NullString
		var ingredientQuantity sql.NullFloat64
		var ingredientPrice sql.NullFloat64
		var packageId sql.NullInt64
		var metric sql.NullString
		var quantity sql.NullFloat64
		var packagePrice sql.NullFloat64

		err := rows.Scan(&recipeId, &recipeName, &recipePrice, &ingredientId, &ingredientName, &ingredientQuantity, &ingredientPrice, &packageId,
			&metric, &quantity, &packagePrice)

		if pkg.HasError(err) {
			return nil
		}

		var ingredientPackage *models.RecipeIngredientPackage
		var recipeIngredient *models.RecipeIngredient

		if packageId.Valid {
			ingredientPackage = models.NewRecipeIngredientPackage(db.GetLong(packageId), db.GetString(metric), db.GetFloat(quantity),
				db.GetFloat(packagePrice))
		}

		if ingredientId.Valid {
			recipeIngredient = models.NewRecipeIngredient(ingredientId.Int64, ingredientName.String, ingredientPackage,
				ingredientQuantity.Float64, db.GetFloat(ingredientPrice))
		}

		recipe := util.GetValue(recipesGrouper, recipeId)
		if recipe == nil {
			ingredients := util.NewList[models.RecipeIngredient]()
			recipe = models.NewRecipe(recipeId, recipeName, ingredients, &recipePrice)
		}

		if recipeIngredient != nil {
			util.Add(&recipe.Ingredients, *recipeIngredient)
		}
		util.PutValue(&recipesGrouper, recipeId, *recipe)
	}

	recipes := util.NewList[models.Recipe]()

	for _, recipe := range recipesGrouper {
		util.Add(&recipes, recipe)
	}

	return &recipes
}

func ToRecipeDTOList(recipes *[]models.Recipe) *[]dto.RecipeDTO {
	dtos := util.NewList[dto.RecipeDTO]()

	for _, recipe := range *recipes {
		dto := dto.RecipeDTO{
			Id:          recipe.Id,
			Name:        recipe.Name,
			Price:       *recipe.Price,
			Ingredients: *ToRecipeIngredientDTOList(&recipe.Ingredients),
		}

		util.Add(&dtos, dto)
	}

	return &dtos
}
