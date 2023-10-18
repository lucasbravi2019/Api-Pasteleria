package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

func ToIngredientList(rows *sql.Rows) (*[]models.Ingredient, error) {
	ingredients := util.NewList[models.Ingredient]()

	for rows.Next() {
		var id int64
		var name string

		err := rows.Scan(&id, &name)

		if err != nil {
			return nil, err
		}

		ingredient := models.NewIngredient(id, name)

		util.Add(&ingredients, *ingredient)
	}
	return &ingredients, nil
}

func ToIngredientDTOList(ingredients []models.Ingredient) (*[]dto.IngredientDTO, error) {
	dtos := util.NewList[dto.IngredientDTO]()

	for _, ingredient := range ingredients {
		dto := dto.NewIngredientDTO(ingredient.Id, ingredient.Name)

		util.Add(&dtos, *dto)
	}

	return &dtos, nil
}

func ToIngredientPackageDTOList(rows *sql.Rows) (*[]dto.IngredientDTO, error) {
	ingredientsById := util.NewMap[int64, dto.IngredientDTO]()
	ingredients := util.NewList[dto.IngredientDTO]()

	for rows.Next() {
		var id int64
		var name string
		var packageId int64
		var price float64
		var metric string
		var quantity float64

		err := rows.Scan(&id, &name, &packageId, &price, &metric, &quantity)

		if err != nil {
			return nil, err
		}

		pkg := dto.NewPackageDTO(packageId, metric, quantity).SetPrice(price)

		ingredientExisting := util.GetValue(ingredientsById, id)

		if ingredientExisting == nil {
			ingredientExisting = dto.NewIngredientDTO(id, name)
		}

		ingredientExisting.AddPackage(pkg)

		util.PutValue(&ingredientsById, id, *ingredientExisting)
	}

	for _, ingredient := range ingredientsById {
		util.Add[dto.IngredientDTO](&ingredients, ingredient)
	}

	return &ingredients, nil
}
