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

	var recipe *Recipe = &Recipe{}

	err := r.db.FindOne(ctx, GetRecipeById(*oid)).Decode(recipe)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, recipe
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

	recipe.ID = *oid

	err := r.db.FindOneAndUpdate(ctx, GetRecipeById(*oid), UpdateRecipe(*recipe)).Decode(recipe)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, recipe
}

func (r *recipeRepository) DeleteRecipe(oid *primitive.ObjectID) (int, *Recipe) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var recipeDeleted *Recipe = &Recipe{}

	err := r.db.FindOneAndDelete(ctx, GetRecipeById(*oid)).Decode(recipeDeleted)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, recipeDeleted
}
