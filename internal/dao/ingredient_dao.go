package dao

import (
	"database/sql"
	"log"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type IngredientDao struct {
	DB *sql.DB
}

type IngredientDaoInterface interface {
	GetAllIngredients() (*[]dto.IngredientDTO, error)
	ValidateExistingIngredient(ingredientName *dto.IngredientNameDTO) error
	CreateIngredient(ingredientName *dto.IngredientNameDTO) error
	UpdateIngredient(id *int64, dto *dto.IngredientNameDTO) error
	DeleteIngredient(id *int64) error
}

var IngredientDaoInstance *IngredientDao

func (d *IngredientDao) GetAllIngredients() (*[]dto.IngredientDTO, error) {
	query, err := db.GetQueryByName(db.Ingredient_FindAll)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query)

	if pkg.HasError(err) {
		log.Println(err.Error())
	}

	ingredients, err := mapper.ToIngredientList(rows)

	if pkg.HasError(err) {
		return nil, err
	}

	dtos, err := mapper.ToIngredientDTOList(*ingredients)

	if pkg.HasError(err) {
		return nil, err
	}

	return dtos, nil
}

func (d *IngredientDao) CreateIngredient(ingredientName *dto.IngredientNameDTO) error {
	query, err := db.GetQueryByName(db.Ingredient_Create)

	if pkg.HasError(err) {
		return err
	}
	_, err = d.DB.Exec(query, ingredientName.Name)

	return err
}

func (d *IngredientDao) UpdateIngredient(id int64, dto *dto.IngredientNameDTO) error {
	query, err := db.GetQueryByName(db.Ingredient_UpdateById)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, dto.Name, id)

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
