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
	FindRecipeByOID(oid *primitive.ObjectID) (*dto.RecipeDTO, error)
	FindRecipesByPackageId(oid *primitive.ObjectID) ([]dto.RecipeDTO, error)
	CreateRecipe(recipe *models.Recipe) (*primitive.ObjectID, error)
	UpdateRecipeName(oid *primitive.ObjectID, recipeName *dto.RecipeNameDTO) error
	DeleteRecipe(oid *primitive.ObjectID) error
	UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error
	UpdateRecipesPrice() error
	GetRecipesByIngredientId(oid *primitive.ObjectID) (*[]models.Recipe, error)
	UpdateRecipes(recipes *[]models.Recipe) error
}

var RecipeDaoInstance *RecipeDao

func (d *RecipeDao) FindAllRecipes() (*[]dto.RecipeDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, queries.All())

	recipes := &[]dto.RecipeDTO{}

	if err != nil {
		log.Println(err.Error())
		return recipes, err
	}

	err = cursor.All(ctx, recipes)

	if err != nil {
		log.Println(err.Error())
		return recipes, err
	}

	return recipes, nil
}

func (d *RecipeDao) FindRecipeByOID(oid *primitive.ObjectID) (*dto.RecipeDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	recipe := &dto.RecipeDTO{}

	err := d.DB.FindOne(ctx, queries.GetRecipeById(*oid)).Decode(recipe)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return recipe, nil
}

func (d *RecipeDao) FindRecipesByPackageId(packageId *primitive.ObjectID) ([]dto.RecipeDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := d.DB.Find(ctx, queries.GetRecipeByPackageId(*packageId))

	recipes := &[]dto.RecipeDTO{}

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	err = cursor.All(ctx, recipes)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return *recipes, nil
}

func (d *RecipeDao) CreateRecipe(recipe *models.Recipe) (*primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := d.DB.InsertOne(ctx, recipe)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id, nil
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

func (d *RecipeDao) GetRecipesByIngredientId(oid *primitive.ObjectID) (*[]models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cur, err := d.DB.Find(ctx, queries.GetRecipeByIngredientId(*oid))

	if err != nil {
		return nil, err
	}

	recipes := &[]models.Recipe{}

	err = cur.All(ctx, recipes)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (d *RecipeDao) UpdateRecipes(recipes []models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	for i := 0; i < len(recipes); i++ {
		_, err := d.DB.UpdateOne(ctx, queries.GetRecipeById(recipes[i].ID), queries.UpdateRecipe(recipes[i]))

		if err != nil {
			return err
		}
	}

	return nil
}
