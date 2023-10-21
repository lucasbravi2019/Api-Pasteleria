package mapper

import (
	"database/sql"
	"log"

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
		var ingredientId int64
		var ingredientName string
		var ingredientQuantity float64
		var ingredientPrice float64
		var packageId int64
		var metric string
		var quantity float64
		var packagePrice float64

		err := rows.Scan(&recipeId, &recipeName, &recipePrice, &ingredientId, &ingredientName, &ingredientQuantity, &ingredientPrice, &packageId,
			&metric, &quantity, &packagePrice)

		if pkg.HasError(err) {
			log.Println(err)
			return nil
		}

		ingredientPackage := models.NewRecipeIngredientPackage(packageId, metric, quantity, packagePrice)
		recipeIngredient := models.NewRecipeIngredient(ingredientId, ingredientName, *ingredientPackage, ingredientQuantity, ingredientPrice)

		recipe := util.GetValue(recipesGrouper, recipeId)
		if recipe == nil {
			recipe = models.NewRecipe(recipeId, recipeName, util.NewList[models.RecipeIngredient](), recipePrice)
		}

		util.Add(&recipe.Ingredients, *recipeIngredient)
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
			Price:       recipe.Price,
			Ingredients: *ToRecipeIngredientDTOList(&recipe.Ingredients),
		}

		util.Add(&dtos, dto)
	}

	return &dtos
}
