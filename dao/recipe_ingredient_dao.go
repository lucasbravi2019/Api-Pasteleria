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
	RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *models.Recipe) error
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

<<<<<<< HEAD
func (d *RecipeIngredientDao) RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *models.RecipeIngredient) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
=======
func (d *RecipeIngredientDao) RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
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
