package recipes

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type recipeRepository struct {
	db *mongo.Collection
}

type RecipeRepository interface {
	FindAllRecipes() (int, *[]RecipeDTO)
	FindRecipeByOID(oid *primitive.ObjectID) (int, *RecipeDTO)
	CreateRecipe(recipe *RecipeNameDTO) (int, *RecipeDTO)
	UpdateRecipeName(oid *primitive.ObjectID, recipeName *RecipeNameDTO) (int, *RecipeDTO)
	AddIngredientToRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) (int, *RecipeDTO)
	DeleteRecipe(oid *primitive.ObjectID) (int, *primitive.ObjectID)
}

var recipeRepositoryInstance *recipeRepository

func (r *recipeRepository) FindAllRecipes() (int, *[]RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := r.db.Aggregate(ctx, GetAggregateAllRecipe())

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	var recipes *[]RecipeDTO = &[]RecipeDTO{}
	err = cursor.All(ctx, recipes)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, recipes
}

func (r *recipeRepository) FindRecipeByOID(oid *primitive.ObjectID) (int, *RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var recipe []RecipeDTO = []RecipeDTO{}

	cursor, err := r.db.Aggregate(ctx, GetAggregateRecipeById(*oid))

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	err = cursor.All(ctx, &recipe)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	if len(recipe) == 0 {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, &recipe[0]
}

func (r *recipeRepository) CreateRecipe(recipe *RecipeNameDTO) (int, *RecipeDTO) {
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

	var recipeCreated *RecipeDTO = &RecipeDTO{
		ID:   oid.(primitive.ObjectID),
		Name: recipe.Name,
	}

	return http.StatusCreated, recipeCreated
}

func (r *recipeRepository) UpdateRecipeName(oid *primitive.ObjectID, recipeName *RecipeNameDTO) (int, *RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*oid), UpdateRecipeName(*recipeName))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	var recipe *RecipeDTO = &RecipeDTO{}

	cursor, err := r.db.Aggregate(ctx, GetAggregateRecipeById(*oid))

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	err = cursor.Decode(recipe)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipe
}

func (r *recipeRepository) AddIngredientToRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) (int, *RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*oid), AddIngredientToRecipe(*recipe))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	var dto []RecipeDTO = []RecipeDTO{}

	cursor, err := r.db.Aggregate(ctx, GetAggregateRecipeById(*oid))

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	err = cursor.All(ctx, &dto)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	if len(dto) == 0 {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, &dto[0]
}

func (r *recipeRepository) DeleteRecipe(oid *primitive.ObjectID) (int, *primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.DeleteOne(ctx, GetRecipeById(*oid))

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, oid
}
