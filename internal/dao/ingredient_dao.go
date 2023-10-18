package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ChangeIngredientPrice(packageId *int64, priceDTO *dto.IngredientPackagePriceDTO) error
}

var IngredientDaoInstance *IngredientDao

func (d *IngredientDao) GetAllIngredients() (*[]dto.IngredientDTO, error) {
	query, err := db.GetQueryByName(db.Ingredient_FindAll)

	if err != nil {
		return nil, err
	}

	rows, err := d.DB.Query(query)

	if err != nil {
		log.Println(err.Error())
	}

	ingredients, err := mapper.ToIngredientList(rows)

	if err != nil {
		return nil, err
	}

	dtos, err := mapper.ToIngredientDTOList(*ingredients)

	if err != nil {
		return nil, err
	}

	return dtos, nil
}

func (d *IngredientDao) CreateIngredient(ingredientName *dto.IngredientNameDTO) error {
	query, err := db.GetQueryByName(db.Ingredient_Create)

	if err != nil {
		return err
	}
	_, err = d.DB.Exec(query, ingredientName.Name)

	return err
}

func (d *IngredientDao) UpdateIngredient(id int64, dto *dto.IngredientNameDTO) error {
	query, err := db.GetQueryByName(db.Ingredient_UpdateById)

	if err != nil {
		return err
	}

	_, err = d.DB.Exec(query, dto.Name, id)

	return err
}

func (d *IngredientDao) DeleteIngredient(id *int64) error {
	query, err := db.GetQueryByName(db.Ingredient_DeleteById)

	if err != nil {
		return err
	}

	_, err = d.DB.Exec(query, id)

	return err
}

func (d *IngredientDao) ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *dto.IngredientPackagePriceDTO) error {

	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
