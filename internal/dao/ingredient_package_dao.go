package dao

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type IngredientPackageDao struct {
	DB *sql.DB
}

type IngredientPackageDaoInterface interface {
	FindAllIngredientPackages() (*[]dto.IngredientDTO, error)
}

var IngredientPackageDaoInstance *IngredientPackageDao

func (d *IngredientPackageDao) FindAllIngredientPackages(id int64) (*[]dto.IngredientDTO, error) {
	query, err := db.GetQueryByName(db.Ingredient_FindAllIngredientPackages)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query, id)

	if pkg.HasError(err) {
		return nil, err
	}

	dtos, err := mapper.ToIngredientPackageDTOList(rows)

	if pkg.HasError(err) {
		return nil, err
	}

	return dtos, nil
}
