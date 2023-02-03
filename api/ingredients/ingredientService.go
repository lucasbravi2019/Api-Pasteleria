package ingredients

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ingredientService struct {
	repository IngredientRepository
}

type IngredientService interface {
	GetAllIngredients() []Ingredient
	CreateIngredient(ingredient Ingredient) (string, error)
	UpdateIngredient(oid primitive.ObjectID, ingredient Ingredient) (string, error)
	DeleteIngredient(oid primitive.ObjectID) error
}

var ingredientServiceInstance *ingredientService

func (s *ingredientService) GetAllIngredients() []Ingredient {
	return s.repository.GetAllIngredients()
}

func (s *ingredientService) CreateIngredient(ingredient Ingredient) (string, error) {
	return s.repository.CreateIngredient(ingredient)
}

func (s *ingredientService) UpdateIngredient(oid primitive.ObjectID, ingredient Ingredient) (string, error) {
	return s.repository.UpdateIngredient(oid, ingredient)
}

func (s *ingredientService) DeleteIngredient(oid primitive.ObjectID) error {
	return s.repository.DeleteIngredient(oid)
}
