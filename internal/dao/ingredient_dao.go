package dao

import (
	"context"
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type IngredientDao struct {
	DB               *sql.DB
	IngredientMapper *mapper.IngredientMapper
}

var IngredientDaoInstance *IngredientDao

func (d *IngredientDao) GetAllIngredients() (*[]dto.IngredientResponse, error) {
	query, err := db.GetQueryByName(db.Ingredient_FindAll)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query)

	if pkg.HasError(err) {
		return nil, err
	}

	ingredients, err := d.IngredientMapper.ToIngredientList(rows)

	if pkg.HasError(err) {
		return nil, err
	}

	return ingredients, nil
}

func (d *IngredientDao) FindIngredientIdByName(ingredientName *string) (*int64, error) {
	query, err := db.GetQueryByName(db.Ingredient_FindIngredientIdByName)

	if pkg.HasError(err) {
		return nil, err
	}

	row := d.DB.QueryRow(query, ingredientName)

	return d.IngredientMapper.ToIngredientId(row)
}

func (d *IngredientDao) CreateIngredientName(ingredientName *string) error {
	query, err := db.GetQueryByName(db.Ingredient_Create)

	if pkg.HasError(err) {
		return err
	}
	_, err = d.DB.Exec(query, ingredientName)

	return err
}

func (d *IngredientDao) UpdateIngredientName(dto *dto.IngredientRequest) error {
	query, err := db.GetQueryByName(db.Ingredient_UpdateById)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, dto.Name, dto.Id)

	return err
}

func (d *IngredientDao) DeleteIngredient(id *int64) error {
	query, err := db.GetQueryByName(db.Ingredient_DeleteById)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, id)

	return err
}

func (d *IngredientDao) AddIngredientPackage(ingredientId *int64, packages *[]dto.IngredientPackageRequest) error {
	query, err := db.GetQueryByName(db.Ingredient_AddPackage)

	if pkg.HasError(err) {
		return err
	}

	tx, err := d.DB.BeginTx(context.TODO(), nil)

	if pkg.HasError(err) {
		return err
	}

	defer func() {
		tx.Commit()
	}()

	for _, newPkg := range *packages {
		_, err := d.DB.Exec(query, ingredientId, &newPkg.Id, &newPkg.Price)

		if pkg.HasError(err) {
			return err
		}
	}

	return nil
}

func (d *IngredientDao) RemoveIngredientPackages(ingredientId *int64) error {
	query, err := db.GetQueryByName(db.Ingredient_DeletePackage)

	if pkg.HasError(err) {
		return err
	}

	tx, err := d.DB.BeginTx(context.TODO(), nil)

	if pkg.HasError(err) {
		return err
	}

	defer func() {
		tx.Commit()
	}()

	_, err = d.DB.Exec(query, ingredientId)

	if pkg.HasError(err) {
		return err
	}

	return nil
}
