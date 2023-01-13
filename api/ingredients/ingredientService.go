package ingredients

import "go.mongodb.org/mongo-driver/bson/primitive"

type ingredientService struct {
	repository IngredientRepository
}

type IngredientService interface {
	GetAllIngredients() []Ingredient
	CreateIngredient(ingredient Ingredient) string
	UpdateIngredient(oid primitive.ObjectID, ingredient Ingredient) string
	DeleteIngredient(oid primitive.ObjectID) string
}

var ingredientServiceInstance *ingredientService

func (s *ingredientService) GetAllIngredients() []Ingredient {
	return s.repository.GetAllIngredients()
}

func (s *ingredientService) CreateIngredient(ingredient Ingredient) string {
	return s.repository.CreateIngredient(ingredient)
}

func (s *ingredientService) UpdateIngredient(oid primitive.ObjectID, ingredient Ingredient) string {
	return s.repository.UpdateIngredient(oid, ingredient)
}

func (s *ingredientService) DeleteIngredient(oid primitive.ObjectID) string {
	return s.repository.DeleteIngredient(oid)
}
