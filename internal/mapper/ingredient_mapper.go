package mapper

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientMapper struct {
	PackageMapper *PackageMapper
}

var IngredientMapperInstance *IngredientMapper

func (m *IngredientMapper) ToIngredientList(rows *sql.Rows) (*[]dto.IngredientResponse, error) {
	ingredientsMap := util.NewMap[int64, dto.IngredientResponse]()

	for rows.Next() {
		var ingredientId int64
		var ingredientName string
		var ingredientPackageId sql.NullInt64
		var ingredientPackagePrice sql.NullFloat64
		var packageId sql.NullInt64
		var metric sql.NullString
		var quantity sql.NullFloat64

		err := rows.Scan(&ingredientId, &ingredientName, &ingredientPackageId, &ingredientPackagePrice, &packageId, &metric, &quantity)

		if pkg.HasError(err) {
			return nil, err
		}

		pkg := m.PackageMapper.ToPackageNullable(packageId, metric, quantity)

		ingredient := util.GetValue(ingredientsMap, ingredientId)

		if ingredient == nil {
			ingredient = m.ToIngredientResponse(&ingredientId, &ingredientName)
		}

		ingredientPackage := m.ToIngredientPackageResponse(ingredientPackageId, ingredientPackagePrice, pkg)
		if ingredientPackage != nil {
			util.Add(ingredient.Packages, *ingredientPackage)
		}

		util.PutValue(&ingredientsMap, &ingredientId, ingredient)
	}

	ingredients := util.NewList[dto.IngredientResponse]()

	for _, ingredient := range ingredientsMap {
		util.Add(&ingredients, ingredient)
	}

	return &ingredients, nil
}

func (m *IngredientMapper) ToIngredientId(row *sql.Row) (*int64, error) {
	var ingredientId sql.NullInt64
	err := row.Scan(&ingredientId)

	if pkg.HasError(err) {
		return nil, err
	}

	return db.GetLong(ingredientId), nil
}

func (m *IngredientMapper) ToIngredient(id *int64, name *string) *dto.Ingredient {
	return &dto.Ingredient{
		Id:   id,
		Name: name,
	}
}

func (m *IngredientMapper) ToIngredientNullable(id sql.NullInt64, name sql.NullString) *dto.Ingredient {
	if !id.Valid {
		return nil
	}

	return &dto.Ingredient{
		Id:   db.GetLong(id),
		Name: db.GetString(name),
	}
}

func (m *IngredientMapper) ToIngredientPackageNullable(id sql.NullInt64, price sql.NullFloat64, ingredient *dto.Ingredient,
	pkg *dto.Package) *dto.IngredientPackage {
	if !id.Valid {
		return nil
	}

	return &dto.IngredientPackage{
		Id:         db.GetLong(id),
		Price:      db.GetFloat(price),
		Ingredient: ingredient,
		Package:    pkg,
	}
}

func (m *IngredientMapper) ToIngredientResponse(id *int64, name *string) *dto.IngredientResponse {
	return &dto.IngredientResponse{
		Id:       id,
		Name:     name,
		Packages: &[]dto.IngredientPackageResponse{},
	}
}

func (m *IngredientMapper) ToIngredientPackageResponse(ingredientPackageId sql.NullInt64, ingredientPackagePrice sql.NullFloat64,
	pkg *dto.Package) *dto.IngredientPackageResponse {
	if !ingredientPackageId.Valid {
		return nil
	}

	return &dto.IngredientPackageResponse{
		IngredientPackageId: db.GetLong(ingredientPackageId),
		Price:               db.GetFloat(ingredientPackagePrice),
		Package:             pkg,
	}
}
