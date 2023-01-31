package ingredients

import (
	"context"
	"errors"
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
	FindIngredientByOID(oid primitive.ObjectID) (Ingredient, error)
	CreateIngredient(ingredient Ingredient) (string, error)
	UpdateIngredient(oid primitive.ObjectID, ingredient Ingredient) (string, error)
	DeleteIngredient(oid primitive.ObjectID) error
}

var ingredientRepositoryInstance *ingredientRepository

func (r *ingredientRepository) GetAllIngredients() []Ingredient {
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
		return []Ingredient{}
	}

	return ingredients
}

func (r *ingredientRepository) FindIngredientByOID(oid primitive.ObjectID) (Ingredient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	result := r.db.FindOne(ctx, filter)

	if result.Err() != nil {
		log.Println(result.Err())
		return Ingredient{}, errors.New("no se pudo encontrar el ingrediente")
	}

	var ingredient Ingredient

	err := result.Decode(&ingredient)

	if err != nil {
		log.Println(err.Error())
		return Ingredient{}, errors.New("no se pudo decodificar la respuesta")
	}

	return ingredient, nil
}

func (r *ingredientRepository) CreateIngredient(ingredient Ingredient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	oid, err := r.db.InsertOne(ctx, ingredient)

	if err != nil {
		log.Println(err.Error())
		return "", errors.New("no pudo crearse el ingrediente")
	}

	return oid.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *ingredientRepository) UpdateIngredient(oid primitive.ObjectID, ingredient Ingredient) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}
	result, err := r.db.ReplaceOne(ctx, filter, ingredient)

	if err != nil {
		log.Println(err.Error())
		return "", errors.New("no pudo actualizarse el ingrediente")
	}

	if result.ModifiedCount < 1 {
		log.Println("El ingrediente no fue actualizado")
		return "", errors.New("no pudo actualizarse el ingrediente")
	}

	return result.UpsertedID.(primitive.ObjectID).Hex(), nil
}

func (r *ingredientRepository) DeleteIngredient(oid primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	result, err := r.db.DeleteOne(ctx, filter)

	if err != nil {
		log.Println(err.Error())
		return errors.New("no pudo borrarse el ingrediente")
	}

	if result.DeletedCount < 1 {
		log.Println("No pudo borrarse el ingrediente")
		return errors.New("no pudo borrarse el ingrediente")
	}

	return nil
}
