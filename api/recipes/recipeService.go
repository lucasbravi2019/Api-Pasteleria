package recipes

import "go.mongodb.org/mongo-driver/bson/primitive"

type recipeService struct {
	repository RecipeRepository
}

type RecipeService interface {
	GetAllRecipes() []Recipe
	GetRecipe(oid primitive.ObjectID) Recipe
}

var recipeServiceInstance *recipeService

func (s *recipeService) GetAllRecipes() []Recipe {
	return s.repository.FindAllRecipes()
}

func (s *recipeService) GetRecipe(oid primitive.ObjectID) Recipe {
	return s.repository.FindRecipeByOID(oid)
}
