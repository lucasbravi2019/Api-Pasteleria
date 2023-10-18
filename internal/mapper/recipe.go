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
	recipes := util.NewList[models.Recipe]()
	for rows.Next() {
		var id int
		var name string
		var price float64

		err := rows.Scan(&id, &name, &price)

		if pkg.HasError(err) {
			log.Println(err)
			return nil
		}

		recipe := models.Recipe{
			ID:    id,
			Name:  name,
			Price: price,
		}

		util.Add(&recipes, recipe)
	}

	return &recipes
}

func ToRecipe(row *sql.Row) *models.Recipe {
	var id int
	var name string
	var price float64

	err := row.Scan(&id, &name, &price)

	if pkg.HasError(err) {
		log.Println(err)
		return nil
	}
	return &models.Recipe{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

func ToRecipeDTO(recipe models.Recipe) *dto.RecipeDTO {
	return &dto.RecipeDTO{
		ID:    recipe.ID,
		Name:  recipe.Name,
		Price: recipe.Price,
	}
}

func ToRecipeDTOList(recipes *[]models.Recipe) *[]dto.RecipeDTO {
	dtos := util.NewList[dto.RecipeDTO]()

	for _, recipe := range *recipes {
		dto := dto.RecipeDTO{
			ID:    recipe.ID,
			Name:  recipe.Name,
			Price: recipe.Price,
		}

		util.Add(&dtos, dto)
	}

	return &dtos
}
