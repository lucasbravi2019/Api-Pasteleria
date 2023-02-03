package ingredients

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ingredientRepository struct {
	db *mongo.Collection
}

type IngredientRepository interface {
	GetAllIngredients() (int, []Ingredient)
	FindIngredientByOID(oid *primitive.ObjectID) (int, *Ingredient)
	CreateIngredient(ingredient *Ingredient) (int, *Ingredient)
	UpdateIngredient(oid *primitive.ObjectID, ingredient *Ingredient) (int, *Ingredient)
	DeleteIngredient(oid *primitive.ObjectID) (int, *Ingredient)
}

var ingredientRepositoryInstance *ingredientRepository

func (r *ingredientRepository) GetAllIngredients() (int, []Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results, err := r.db.Find(ctx, bson.D{})

	if err != nil {
		log.Println(err.Error())
	}

	var ingredients []Ingredient

	err = results.All(ctx, &ingredients)

	if err != nil {
		log.Println(err.Error())
	}

	if len(ingredients) < 1 {
		return http.StatusOK, []Ingredient{}
	}

	return http.StatusOK, ingredients
}

func (r *ingredientRepository) FindIngredientByOID(oid *primitive.ObjectID) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	result := r.db.FindOne(ctx, filter)

	if result.Err() != nil {
		log.Println(result.Err())
		return http.StatusNotFound, nil
	}

	var ingredient *Ingredient = &Ingredient{}

	err := result.Decode(ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, ingredient
}

func (r *ingredientRepository) CreateIngredient(ingredient *Ingredient) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := r.db.InsertOne(ctx, *ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	oid := result.InsertedID

	if oid == nil {
		return http.StatusInternalServerError, nil
	}

	var ingredientCreated *Ingredient = &Ingredient{
		ID:       oid.(primitive.ObjectID),
		Name:     ingredient.Name,
		Metric:   ingredient.Metric,
		Quantity: ingredient.Quantity,
		Price:    ingredient.Price,
	}

	return http.StatusCreated, ingredientCreated
}

func (r *ingredientRepository) UpdateIngredient(oid *primitive.ObjectID, ingredient *Ingredient) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}
	result, err := r.db.ReplaceOne(ctx, filter, ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	uid := result.UpsertedID

	if uid == nil {
		return http.StatusInternalServerError, nil
	}

	var ingredientUpdated *Ingredient = &Ingredient{
		ID:       uid.(primitive.ObjectID),
		Name:     ingredient.Name,
		Metric:   ingredient.Metric,
		Quantity: ingredient.Quantity,
		Price:    ingredient.Price,
	}

	return http.StatusOK, ingredientUpdated
}

func (r *ingredientRepository) DeleteIngredient(oid *primitive.ObjectID) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	result := r.db.FindOneAndDelete(ctx, filter)

	var ingredientDeleted *Ingredient = &Ingredient{}

	err := result.Decode(ingredientDeleted)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, ingredientDeleted
}
