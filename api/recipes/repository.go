package recipes

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type recipeRepository struct {
	db *mongo.Collection
}

type RecipeRepository interface {
	FindAllRecipes() (int, []Recipe)
	FindRecipeByOID(oid *primitive.ObjectID) (int, *Recipe)
	CreateRecipe(recipe *RecipeName) (int, *Recipe)
	UpdateRecipe(oid *primitive.ObjectID, recipe *Recipe) (int, *Recipe)
	DeleteRecipe(oid *primitive.ObjectID) (int, *Recipe)
}

var recipeRepositoryInstance *recipeRepository

func (r *recipeRepository) FindAllRecipes() (int, []Recipe) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := r.db.Find(ctx, bson.D{})

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	var recipes []Recipe
	err = result.All(ctx, &recipes)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	if len(recipes) < 1 {
		return http.StatusOK, []Recipe{}
	}

	return http.StatusOK, recipes
}

func (r *recipeRepository) FindRecipeByOID(oid *primitive.ObjectID) (int, *Recipe) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	filter := bson.D{primitive.E{Key: "_id", Value: oid}}
	result := r.db.FindOne(ctx, filter)

	var recipe Recipe

	err := result.Decode(&recipe)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, &recipe
}

func (r *recipeRepository) CreateRecipe(recipe *RecipeName) (int, *Recipe) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := r.db.InsertOne(ctx, recipe)
	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	oid := result.InsertedID

	if oid == nil {
		log.Println("No pudo insertarse en la base de datos")
		return http.StatusInternalServerError, nil
	}

	var recipeCreated *Recipe = &Recipe{
		ID:   oid.(primitive.ObjectID),
		Name: recipe.Name,
	}

	return http.StatusCreated, recipeCreated
}

func (r *recipeRepository) UpdateRecipe(oid *primitive.ObjectID, recipe *Recipe) (int, *Recipe) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	recipe.ID = *oid

	result, err := r.db.ReplaceOne(ctx, filter, recipe)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	if result.ModifiedCount < 1 {
		log.Println("No se actualizó la receta")
	}

	var recipeUpdated *Recipe = &Recipe{
		ID:          *oid,
		Name:        recipe.Name,
		Ingredients: recipe.Ingredients,
		Price:       recipe.Price,
	}

	return http.StatusOK, recipeUpdated
}

func (r *recipeRepository) DeleteRecipe(oid *primitive.ObjectID) (int, *Recipe) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.D{primitive.E{Key: "_id", Value: oid}}

	result := r.db.FindOneAndDelete(ctx, filter)

	var recipeDeleted *Recipe = &Recipe{}

	err := result.Decode(recipeDeleted)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	if recipeDeleted.ID.IsZero() {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipeDeleted
}
