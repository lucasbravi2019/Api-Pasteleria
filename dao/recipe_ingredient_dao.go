package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeIngredientDao struct {
	DB *sql.DB
}

type RecipeIngredientDaoInterface interface {
	AddIngredientToRecipe(oid *primitive.ObjectID, recipe *models.RecipeIngredient) error
	RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *models.RecipeIngredient) error
	RemoveIngredientByPackageId(packageId *primitive.ObjectID) error
	UpdateIngredientPackagePrice(packageId *primitive.ObjectID, price float64) error
	UpdateIngredientsPrice(packageId *primitive.ObjectID, recipe dto.RecipeDTO) error
}

var RecipeIngredientDaoInstance *RecipeIngredientDao

func (d *RecipeIngredientDao) AddIngredientToRecipe(oid *primitive.ObjectID, recipe *models.RecipeIngredient) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *models.RecipeIngredient) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) RemoveIngredientByPackageId(packageId *primitive.ObjectID) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) UpdateIngredientPackagePrice(packageId *primitive.ObjectID, price float64) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) UpdateIngredientsPrice(packageId *primitive.ObjectID, recipe dto.RecipeDTO) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
