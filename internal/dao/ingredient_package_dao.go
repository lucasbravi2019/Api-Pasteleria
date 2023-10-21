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
	UpdateIngredientPackages(ingredient *dto.IngredientPackagePrices) error
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

func (d *IngredientPackageDao) UpdateIngredientPackages(ingredient *dto.IngredientPackagePrices) error {
	tx, err := d.DB.Begin()

	if pkg.HasError(err) {
		return err
	}

	defer func() {
		if pkg.HasError(err) {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query, err := db.GetQueryByName(db.Ingredient_DeletePackage)

	if pkg.HasError(err) {
		return err
	}

	_, err = tx.Exec(query, ingredient.IngredientId)

	if pkg.HasError(err) {
		return err
	}

	query, err = db.GetQueryByName(db.Ingredient_AddPackage)

	if pkg.HasError(err) {
		return err
	}

	for _, envase := range ingredient.Packages {
		_, err = tx.Exec(query, ingredient.IngredientId, envase.PackageId, envase.Price)

		if pkg.HasError(err) {
			return err
		}
	}

	return nil
}
