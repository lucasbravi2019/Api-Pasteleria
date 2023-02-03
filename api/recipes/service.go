package recipes

import (
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/ingredients"
	"github.com/lucasbravi2019/pasteleria/core"
)

type recipeService struct {
	recipeRepository     RecipeRepository
	ingredientRepository ingredients.IngredientRepository
}

type RecipeService interface {
	GetAllRecipes() (int, []Recipe)
	GetRecipe(r *http.Request) (int, *Recipe)
	CreateRecipe(r *http.Request) (int, *Recipe)
	UpdateRecipe(r *http.Request) (int, *Recipe)
	DeleteRecipe(r *http.Request) (int, *Recipe)
	AddIngredientToRecipe(r *http.Request) (int, *Recipe)
}

var recipeServiceInstance *recipeService

func (s *recipeService) GetAllRecipes() (int, []Recipe) {
	return s.recipeRepository.FindAllRecipes()
}

func (s *recipeService) GetRecipe(r *http.Request) (int, *Recipe) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.recipeRepository.FindRecipeByOID(oid)
}

func (s *recipeService) CreateRecipe(r *http.Request) (int, *Recipe) {
	var recipeName *RecipeName = &RecipeName{}

	invalidBody := core.DecodeBody(r, recipeName)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.recipeRepository.CreateRecipe(recipeName)
}

func (s *recipeService) UpdateRecipe(r *http.Request) (int, *Recipe) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	var recipe *Recipe = &Recipe{}

	invalidBody := core.DecodeBody(r, recipe)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.recipeRepository.UpdateRecipe(oid, recipe)
}

func (s *recipeService) DeleteRecipe(r *http.Request) (int, *Recipe) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.recipeRepository.DeleteRecipe(oid)
}

func (s *recipeService) AddIngredientToRecipe(r *http.Request) (int, *Recipe) {
	recipeOid := core.ConvertHexToObjectId(mux.Vars(r)["recipeId"])

	if recipeOid == nil {
		return http.StatusBadRequest, nil
	}

	_, recipe := s.recipeRepository.FindRecipeByOID(recipeOid)

	if recipe == nil {
		return http.StatusNotFound, nil
	}

	ingredientOid := core.ConvertHexToObjectId(mux.Vars(r)["ingredientId"])

	if ingredientOid == nil {
		return http.StatusBadRequest, nil
	}

	_, ingredient := s.ingredientRepository.FindIngredientByOID(ingredientOid)

	if ingredient == nil {
		return http.StatusNotFound, nil
	}

	var ingredientDetails *IngredientDetails = &IngredientDetails{}

	invalidBody := core.DecodeBody(r, ingredientDetails)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	err := validate(ingredient, ingredientDetails)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	var recipeIngredient *RecipeIngredient = &RecipeIngredient{
		Ingredient: *ingredient,
		Price:      float32(ingredientDetails.Quantity) / float32(ingredient.Quantity) * ingredient.Price,
		Quantity:   ingredientDetails.Quantity,
	}

	recipe.Ingredients = append(recipe.Ingredients, *recipeIngredient)
	recipe.Price = func() float32 {
		var result float32
		for _, ingredient := range recipe.Ingredients {
			result += ingredient.Price
		}
		return result
	}()

	_, recipeUpdated := s.recipeRepository.UpdateRecipe(recipeOid, recipe)

	if recipeUpdated == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipeUpdated
}

func validate(ingredient *ingredients.Ingredient, ingredientDetails *IngredientDetails) error {
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
