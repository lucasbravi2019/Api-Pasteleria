package recipes

import "go.mongodb.org/mongo-driver/bson/primitive"

type recipeService struct {
	repository RecipeRepository
}

type RecipeService interface {
	GetAllRecipes() []Recipe
	GetRecipe(oid primitive.ObjectID) Recipe
	CreateRecipe(recipe Recipe) string
	UpdateRecipe(oid primitive.ObjectID, recipe Recipe) string
	DeleteRecipe(oid primitive.ObjectID) string
}

var recipeServiceInstance *recipeService

func (s *recipeService) GetAllRecipes() []Recipe {
	return s.repository.FindAllRecipes()
}

func (s *recipeService) GetRecipe(oid primitive.ObjectID) Recipe {
	return s.repository.FindRecipeByOID(oid)
}

func (s *recipeService) CreateRecipe(recipe Recipe) string {
	return s.repository.CreateRecipe(recipe)
}

func (s *recipeService) UpdateRecipe(oid primitive.ObjectID, recipe Recipe) string {
	return s.repository.UpdateRecipe(oid, recipe)
}

func (s *recipeService) DeleteRecipe(oid primitive.ObjectID) string {
	return s.repository.DeleteRecipe(oid)
}
