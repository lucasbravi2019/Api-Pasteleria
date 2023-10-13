package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"github.com/lucasbravi2019/pasteleria/util"
)

func ToRecipeList(rows *sql.Rows) []models.Recipe {
	recipes := make(map[int]models.Recipe)

	for rows.Next() {
		var id int
		var name string
		var ingredient int
		var price float64

		if err := rows.Scan(&id, &name, &ingredient, &price); err != nil {
			return nil
		}

		recipe := recipes[id]

		if recipe.ID != 0 {
			util.Add(recipe.Ingredients, models.RecipeIngredient{ID: ingredient})
		} else {
			recipes[id] = models.Recipe{
				ID:          id,
				Name:        name,
				Price:       price,
				Ingredients: util.ToList(models.RecipeIngredient{ID: ingredient}),
			}
		}

	}

	var recipesMapped []models.Recipe

	for _, v := range recipes {
		util.Add(recipesMapped, v)
	}

	return recipesMapped
}

func ToRecipe(rows *sql.Rows) *models.Recipe {
	recipes := make(map[int]models.Recipe)

	var recipeId int
	for rows.Next() {
		var id int
		var name string
		var ingredient int
		var price float64

		if err := rows.Scan(&id, &name, &ingredient, &price); err != nil {
			return nil
		}

		recipeId = id
		recipe := recipes[id]

		if recipe.ID != 0 {
			util.Add(recipe.Ingredients, models.RecipeIngredient{ID: ingredient})
		} else {
			recipes[id] = models.Recipe{
				ID:          id,
				Name:        name,
				Price:       price,
				Ingredients: util.ToList(models.RecipeIngredient{ID: ingredient}),
			}
		}

	}

	if recipeId == 0 {
		return nil
	}

	recipe := recipes[recipeId]
	return &recipe
}

func ToRecipeDTO(rows *sql.Rows) *dto.RecipeDTO {
	return nil
}

func ToRecipeDTOList(rows *sql.Rows) *[]dto.RecipeDTO {
	return nil
}
