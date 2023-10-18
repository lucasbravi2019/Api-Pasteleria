package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
)

type RecipeIngredientDao struct {
	DB *sql.DB
}

type RecipeIngredientDaoInterface interface {
	AddIngredientToRecipe(oid *int64, recipe *models.RecipeIngredient) error
	RemoveIngredientFromRecipe(oid *int64, recipe *models.Recipe) error
	RemoveIngredientByPackageId(packageId *int64) error
	UpdateIngredientPackagePrice(packageId *int64, price float64) error
	UpdateIngredientsPrice(packageId *int64, recipe dto.RecipeDTO) error
}

var RecipeIngredientDaoInstance *RecipeIngredientDao

func (d *RecipeIngredientDao) AddIngredientToRecipe(oid *int64, recipe *models.RecipeIngredient) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) RemoveIngredientFromRecipe(oid *int64, recipe *models.RecipeIngredient) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) RemoveIngredientByPackageId(packageId *int64) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) UpdateIngredientPackagePrice(packageId *int64, price float64) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) UpdateIngredientsPrice(packageId *int64, recipe dto.RecipeDTO) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
