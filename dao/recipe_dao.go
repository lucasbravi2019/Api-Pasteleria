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

type RecipeDao struct {
	DB *mongo.Collection
}

type RecipeDaoInterface interface {
	FindAllRecipes() *[]dto.RecipeDTO
	FindRecipeByOID(oid *primitive.ObjectID) *dto.RecipeDTO
	FindRecipesByPackageId(oid *primitive.ObjectID) []dto.RecipeDTO
	CreateRecipe(recipe *models.Recipe) *primitive.ObjectID
	UpdateRecipeName(oid *primitive.ObjectID, recipeName *dto.RecipeNameDTO) error

	DeleteRecipe(oid *primitive.ObjectID) error

	UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error

	UpdateRecipesPrice() error
}

var RecipeDaoInstance *RecipeDao

func (d *RecipeDao) FindAllRecipes() *[]dto.RecipeDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, queries.All())

	recipes := &[]dto.RecipeDTO{}

	if err != nil {
		log.Println(err.Error())
		return recipes
	}

	err = cursor.All(ctx, recipes)

	if err != nil {
		log.Println(err.Error())
	}

	return recipes
}

func (d *RecipeDao) FindRecipeByOID(oid *primitive.ObjectID) *dto.RecipeDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	recipe := &dto.RecipeDTO{}

	err := d.DB.FindOne(ctx, queries.GetRecipeById(*oid)).Decode(recipe)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return recipe
}

func (d *RecipeDao) FindRecipesByPackageId(packageId *primitive.ObjectID) []dto.RecipeDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, queries.GetRecipeByPackageId(*packageId))

	recipes := &[]dto.RecipeDTO{}

	if err != nil {
		log.Println(err.Error())
		return *recipes
	}

	err = cursor.All(ctx, &recipes)

	if err != nil {
		log.Println(err.Error())
	}

	return *recipes
}

func (d *RecipeDao) CreateRecipe(recipe *models.Recipe) *primitive.ObjectID {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := d.DB.InsertOne(ctx, recipe)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if result.InsertedID == nil {
		return nil
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id
}

func (d *RecipeDao) UpdateRecipeName(oid *primitive.ObjectID, recipe *models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateOne(ctx, queries.GetRecipeById(*oid), queries.UpdateRecipeName(*recipe))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeDao) DeleteRecipe(oid *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.DeleteOne(ctx, queries.GetRecipeById(*oid))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeDao) UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateOne(ctx, queries.GetRecipeById(*recipeId), queries.SetRecipePrice())

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeDao) UpdateRecipesPrice() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateMany(ctx, queries.All(), queries.SetRecipePrice())

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
