package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

func ToIngredientList(rows *sql.Rows) (*[]dto.IngredientDTO, error) {
	ingredients := util.NewList[dto.IngredientDTO]()

	for rows.Next() {
		var id int64
		var name string

		err := rows.Scan(&id, &name)

		if pkg.HasError(err) {
			return nil, err
		}

		ingredient := &dto.IngredientDTO{
			Id:   id,
			Name: name,
		}

		util.Add(&ingredients, *ingredient)
	}
	return &ingredients, nil
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

		if pkg.HasError(err) {
			return nil, err
		}

		pkg := &dto.PackageDTO{
			Id:       &packageId,
			Metric:   &metric,
			Quantity: &quantity,
			Price:    &price,
		}

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
