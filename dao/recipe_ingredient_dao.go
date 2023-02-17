package dao

import (
	"context"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"github.com/lucasbravi2019/pasteleria/queries"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecipeIngredientDao struct {
	DB *mongo.Collection
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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateOne(ctx, queries.GetRecipeById(*oid), queries.AddIngredientToRecipe(*recipe))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *models.RecipeIngredient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateOne(ctx, queries.GetRecipeById(*oid), queries.RemoveIngredientFromRecipe(*recipe))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) RemoveIngredientByPackageId(packageId *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateMany(ctx, queries.GetRecipeByPackageId(*packageId), queries.RemovePackageFromRecipes(*packageId))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) UpdateIngredientPackagePrice(packageId *primitive.ObjectID, price float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateMany(ctx, queries.GetRecipeByPackageId(*packageId), queries.SetIngredientPackagePrice(price),
		queries.GetArrayFiltersForIngredientsByPackageId(*packageId))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeIngredientDao) UpdateIngredientsPrice(packageId *primitive.ObjectID, recipe dto.RecipeDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateOne(ctx, queries.GetRecipeByPackageId(*packageId), queries.SetRecipeIngredientPrice(recipe))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
