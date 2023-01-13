package ingredients

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ingredientRepository struct {
	db *mongo.Collection
}

type IngredientRepository interface {
	GetAllIngredients() []Ingredient
	FindIngredientByOID(oid primitive.ObjectID) Ingredient
	CreateIngredient(ingredient Ingredient) string
	UpdateIngredient(oid primitive.ObjectID, ingredient Ingredient) string
	DeleteIngredient(oid primitive.ObjectID) string
}

var ingredientRepositoryInstance *ingredientRepository

func (r *ingredientRepository) GetAllIngredients() []Ingredient {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results, err := r.db.Find(ctx, bson.D{})

	if err != nil {
		log.Fatal(err.Error())
	}

	var ingredients []Ingredient

	err = results.All(ctx, &ingredients)

	if err != nil {
		log.Fatal(err.Error())
	}

	return ingredients
}

func (r *ingredientRepository) FindIngredientByOID(oid primitive.ObjectID) Ingredient {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	result := r.db.FindOne(ctx, filter)

	if result.Err() != nil {
		log.Fatal(result.Err())
	}

	var ingredient Ingredient

	err := result.Decode(&ingredient)

	if err != nil {
		log.Fatal(err.Error())
	}

	return ingredient
}

func (r *ingredientRepository) CreateIngredient(ingredient Ingredient) string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := r.db.InsertOne(ctx, ingredient)

	if err != nil {
		log.Fatal(err.Error())
	}

	return result.InsertedID.(primitive.ObjectID).Hex()
}

func (r *ingredientRepository) UpdateIngredient(oid primitive.ObjectID, ingredient Ingredient) string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}
	result, err := r.db.ReplaceOne(ctx, filter, ingredient)

	if err != nil {
		log.Fatal(err.Error())
	}

	if result.ModifiedCount < 1 {
		log.Println("El registro no fue actualizado")
	}

	return "Se actualizo el ingrediente " + oid.Hex() + " correctamente"
}

func (r *ingredientRepository) DeleteIngredient(oid primitive.ObjectID) string {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	result, err := r.db.DeleteOne(ctx, filter)

	if err != nil {
		log.Fatal(err.Error())
	}

	if result.DeletedCount < 1 {
		log.Fatal("No pudo borrarse el registro")
	}

	return "Se borro el ingrediente " + oid.Hex() + " correctamente"
}
