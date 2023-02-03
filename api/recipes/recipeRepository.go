package recipes

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type recipeRepository struct {
	db *mongo.Collection
}

type RecipeRepository interface {
	FindAllRecipes() ([]Recipe, error)
	FindRecipeByOID(oid primitive.ObjectID) (Recipe, error)
	CreateRecipe(recipe RecipeName) (string, error)
	UpdateRecipe(oid primitive.ObjectID, recipe Recipe) (string, error)
	DeleteRecipe(oid primitive.ObjectID) error
}

var recipeRepositoryInstance *recipeRepository

func (r *recipeRepository) FindAllRecipes() ([]Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := r.db.Find(ctx, bson.D{})

	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("no pudieron encontrarse registros")
	}

	var recipes []Recipe
	err = result.All(ctx, &recipes)

	if err != nil {
		log.Println(err.Error())
		return nil, errors.New("no pudieron encontrarse registros")
	}

	if len(recipes) < 1 {
		return []Recipe{}, nil
	}

	return recipes, nil
}

func (r *recipeRepository) FindRecipeByOID(oid primitive.ObjectID) (Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	result := r.db.FindOne(ctx, filter)

	var recipe Recipe

	err := result.Decode(&recipe)

	if err != nil {
		log.Println(err.Error())
		return Recipe{}, err
	}

	return recipe, nil
}

func (r *recipeRepository) CreateRecipe(recipe RecipeName) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := r.db.InsertOne(ctx, recipe)
	if err != nil {
		log.Println(err.Error())
		return "", errors.New("no pudo crearse la receta")
	}

	oid := result.InsertedID

	if oid == nil {
		log.Println("No pudo insertarse en la base de datos")
		return "", errors.New("no pudo crearse la receta")
	}

	return oid.(primitive.ObjectID).Hex(), nil
}

func (r *recipeRepository) UpdateRecipe(oid primitive.ObjectID, recipe Recipe) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	recipe.ID = oid

	result, err := r.db.ReplaceOne(ctx, filter, recipe)

	if err != nil {
		log.Println(err.Error())
		return "", errors.New("no pudo actualizarse la receta")
	}

	if result.ModifiedCount < 1 {
		log.Println("No se actualizó el registro")
		return "", errors.New("no pudo actualizarse la receta")
	}

	return oid.String(), nil
}

func (r *recipeRepository) DeleteRecipe(oid primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "_id", Value: oid}}

	result, err := r.db.DeleteOne(ctx, filter)

	if err != nil {
		log.Println(err.Error())
		return errors.New("no pudo borrarse receta")
	}

	if result.DeletedCount < 1 {
		log.Println("No pudo borrarse el registro")
		return errors.New("no pudo borrarse receta")
	}

	return nil
}
