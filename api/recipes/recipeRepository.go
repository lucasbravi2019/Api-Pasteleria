package recipes

import (
	"context"
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
	FindAllRecipes() []Recipe
	FindRecipeByOID(oid primitive.ObjectID) Recipe
	CreateRecipe(recipe Recipe) string
	UpdateRecipe(oid primitive.ObjectID, recipe Recipe) string
	DeleteRecipe(oid primitive.ObjectID) string
}

var recipeRepositoryInstance *recipeRepository

func (r *recipeRepository) FindAllRecipes() []Recipe {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := r.db.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var recipes []Recipe
	err = result.All(ctx, &recipes)

	if err != nil {
		log.Fatal(err.Error())
	}

	return recipes
}

func (r *recipeRepository) FindRecipeByOID(oid primitive.ObjectID) Recipe {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	result := r.db.FindOne(ctx, filter)

	var recipe Recipe

	err := result.Decode(&recipe)

	if err != nil {
		log.Fatal(err.Error())
	}

	return recipe
}

func (r *recipeRepository) CreateRecipe(recipe Recipe) string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := r.db.InsertOne(ctx, recipe)
	if err != nil {
		log.Fatal(err.Error())
	}

	oid := result.InsertedID

	if oid == "" {
		log.Fatal("No pudo insertarse en la base de datos")
	}

	return oid.(primitive.ObjectID).Hex()
}

func (r *recipeRepository) UpdateRecipe(oid primitive.ObjectID, recipe Recipe) string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	recipe.ID = oid

	result, err := r.db.ReplaceOne(ctx, filter, recipe)

	if err != nil {
		log.Fatal(err.Error())
	}

	if result.ModifiedCount < 1 {
		log.Println("No se actualizó el registro")
	}

	return oid.Hex()
}

func (r *recipeRepository) DeleteRecipe(oid primitive.ObjectID) string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "_id", Value: oid}}

	result, err := r.db.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err.Error())
	}

	if result.DeletedCount < 1 {
		log.Fatal("No pudo borrarse el registro")
	}

	return "La receta " + oid.Hex() + " pudo borrarse con exito"
}
