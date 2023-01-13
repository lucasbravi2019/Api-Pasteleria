package recipes

import (
	"log"

	"github.com/lucasbravi2019/pasteleria/api/ingredients"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type recipeService struct {
	recipeRepository     RecipeRepository
	ingredientRepository ingredients.IngredientRepository
}

type RecipeService interface {
	GetAllRecipes() []Recipe
	GetRecipe(oid primitive.ObjectID) Recipe
	CreateRecipe(recipe RecipeName) string
	UpdateRecipe(oid primitive.ObjectID, recipe Recipe) string
	DeleteRecipe(oid primitive.ObjectID) string
	AddIngredientToRecipe(recipeOid primitive.ObjectID, ingredientOid primitive.ObjectID, ingredientDetails IngredientDetails) string
}

var recipeServiceInstance *recipeService

func (s *recipeService) GetAllRecipes() []Recipe {
	return s.recipeRepository.FindAllRecipes()
}

func (s *recipeService) GetRecipe(oid primitive.ObjectID) Recipe {
	return s.recipeRepository.FindRecipeByOID(oid)
}

func (s *recipeService) CreateRecipe(recipe RecipeName) string {
	return s.recipeRepository.CreateRecipe(recipe)
}

func (s *recipeService) UpdateRecipe(oid primitive.ObjectID, recipe Recipe) string {
	return s.recipeRepository.UpdateRecipe(oid, recipe)
}

func (s *recipeService) DeleteRecipe(oid primitive.ObjectID) string {
	return s.recipeRepository.DeleteRecipe(oid)
}

func (s *recipeService) AddIngredientToRecipe(recipeOid primitive.ObjectID, ingredientOid primitive.ObjectID,
	ingredientDetails IngredientDetails) string {
	recipe := s.recipeRepository.FindRecipeByOID(recipeOid)

	ingredient := s.ingredientRepository.FindIngredientByOID(ingredientOid)

	if ingredient.Metric != ingredientDetails.Metric {
		log.Println("La unidad de medida no coincide")
		return "La unidad de medida no coincide"
	}

	var recipeIngredient RecipeIngredient = RecipeIngredient{
		Ingredient: ingredient,
		Price:      ingredient.Price * float32(ingredientDetails.Quantity),
		Quantity:   ingredientDetails.Quantity,
	}

	recipe.Ingredients = append(recipe.Ingredients, recipeIngredient)
	recipe.Price = func() float32 {
		var result float32
		for _, ingredient := range recipe.Ingredients {
			result += ingredient.Price
		}
		return result
	}()

	return s.recipeRepository.UpdateRecipe(recipeOid, recipe)
}
