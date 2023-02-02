package recipes

import (
	"errors"
	"log"
	"net/http"

	"github.com/lucasbravi2019/pasteleria/api/ingredients"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type recipeService struct {
	recipeRepository     RecipeRepository
	ingredientRepository ingredients.IngredientRepository
}

type RecipeService interface {
	GetAllRecipes() ([]Recipe, error)
	GetRecipe(oid primitive.ObjectID) (Recipe, error)
	CreateRecipe(recipe RecipeName) (string, error)
	UpdateRecipe(oid primitive.ObjectID, recipe Recipe) (string, error)
	DeleteRecipe(oid primitive.ObjectID) error
	AddIngredientToRecipe(recipeOid primitive.ObjectID, ingredientOid primitive.ObjectID, ingredientDetails IngredientDetails) (error, int, Recipe)
}

var recipeServiceInstance *recipeService

func (s *recipeService) GetAllRecipes() ([]Recipe, error) {
	return s.recipeRepository.FindAllRecipes()
}

func (s *recipeService) GetRecipe(oid primitive.ObjectID) (Recipe, error) {
	return s.recipeRepository.FindRecipeByOID(oid)
}

func (s *recipeService) CreateRecipe(recipe RecipeName) (string, error) {
	return s.recipeRepository.CreateRecipe(recipe)
}

func (s *recipeService) UpdateRecipe(oid primitive.ObjectID, recipe Recipe) (string, error) {
	return s.recipeRepository.UpdateRecipe(oid, recipe)
}

func (s *recipeService) DeleteRecipe(oid primitive.ObjectID) error {
	return s.recipeRepository.DeleteRecipe(oid)
}

func (s *recipeService) AddIngredientToRecipe(recipeOid primitive.ObjectID, ingredientOid primitive.ObjectID,
	ingredientDetails IngredientDetails) (error, int, Recipe) {
	recipe, err := s.recipeRepository.FindRecipeByOID(recipeOid)

	if err != nil {
		return err, http.StatusNotFound, Recipe{}
	}

	ingredient, err := s.ingredientRepository.FindIngredientByOID(ingredientOid)

	if err != nil {
		return errors.New("no se encontró el ingrediente"), http.StatusNotFound, Recipe{}
	}

	err = validate(ingredient, ingredientDetails)

	if err != nil {
		return err, http.StatusBadRequest, Recipe{}
	}

	var recipeIngredient RecipeIngredient = RecipeIngredient{
		Ingredient: ingredient,
		Price:      float32(ingredientDetails.Quantity) / float32(ingredient.Quantity) * ingredient.Price,
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

	_, err = s.recipeRepository.UpdateRecipe(recipeOid, recipe)

	if err != nil {
		return err, http.StatusInternalServerError, Recipe{}
	}

	return nil, http.StatusOK, recipe
}

func validate(ingredient ingredients.Ingredient, ingredientDetails IngredientDetails) error {
	if ingredient.Metric != ingredientDetails.Metric {
		log.Println("La unidad de medida no coincide")
		return errors.New("la unidad de medida no coincide")
	}

	if ingredientDetails.Quantity == 0 {
		log.Println("La cantidad del ingrediente no puede ser 0")
		return errors.New("la cantidad del ingrediente no puede ser 0")
	}
	return nil
}
