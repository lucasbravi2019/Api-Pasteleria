package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/models"
)

type IngredientPackageDao struct {
	DB *sql.DB
}

type IngredientPackageDaoInterface interface {
	UpdateIngredientPackagePrice(packageId *int64, price float64) error
	AddPackageToIngredient(ingredientId *int64, packageId *int64, envase *models.IngredientPackage) error
	FindAllIngredientPackages() (*[]dto.IngredientDTO, error)
}

var IngredientPackageDaoInstance *IngredientPackageDao

func (d *IngredientPackageDao) AddPackageToIngredient(ingredientId *int64, packageId *int64, envase *models.IngredientPackage) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *IngredientPackageDao) RemovePackageFromIngredients(dto dto.IngredientDTO) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *IngredientPackageDao) FindAllIngredientPackages(id int64) (*[]dto.IngredientDTO, error) {
	query, err := db.GetQueryByName(db.Ingredient_FindAllIngredientPackages)

	if err != nil {
		return nil, err
	}

	rows, err := d.DB.Query(query, id)

	if err != nil {
		return nil, err
	}

	dtos, err := mapper.ToIngredientPackageDTOList(rows)

	if err != nil {
		return nil, err
	}

	return dtos, nil
}
