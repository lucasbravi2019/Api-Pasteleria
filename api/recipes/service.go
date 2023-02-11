package recipes

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/ingredients"
	"github.com/lucasbravi2019/pasteleria/api/packages"
	"github.com/lucasbravi2019/pasteleria/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type recipeService struct {
	recipeRepository     RecipeRepository
	ingredientRepository ingredients.IngredientRepository
}

type RecipeService interface {
	GetAllRecipes() (int, *[]RecipeDTO)
	GetRecipe(r *http.Request) (int, *RecipeDTO)
	CreateRecipe(r *http.Request) (int, *RecipeDTO)
	UpdateRecipeName(r *http.Request) (int, *RecipeDTO)
	DeleteRecipe(r *http.Request) (int, *primitive.ObjectID)
	AddIngredientToRecipe(r *http.Request) (int, *RecipeDTO)
}

var recipeServiceInstance *recipeService

func (s *recipeService) GetAllRecipes() (int, *[]RecipeDTO) {
	return s.recipeRepository.FindAllRecipes()
}

func (s *recipeService) GetRecipe(r *http.Request) (int, *RecipeDTO) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.recipeRepository.FindRecipeByOID(oid)
}

func (s *recipeService) CreateRecipe(r *http.Request) (int, *RecipeDTO) {
	var recipeName *RecipeNameDTO = &RecipeNameDTO{}

	invalidBody := core.DecodeBody(r, recipeName)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.recipeRepository.CreateRecipe(recipeName)
}

func (s *recipeService) UpdateRecipeName(r *http.Request) (int, *RecipeDTO) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	var recipe *RecipeNameDTO = &RecipeNameDTO{}

	invalidBody := core.DecodeBody(r, recipe)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.recipeRepository.UpdateRecipeName(oid, recipe)
}

func (s *recipeService) DeleteRecipe(r *http.Request) (int, *primitive.ObjectID) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.recipeRepository.DeleteRecipe(oid)
}

func (s *recipeService) AddIngredientToRecipe(r *http.Request) (int, *RecipeDTO) {
	recipeOid := core.ConvertHexToObjectId(mux.Vars(r)["recipeId"])
	log.Println(recipeOid)
	if recipeOid == nil {
		return http.StatusBadRequest, nil
	}

	_, recipe := s.recipeRepository.FindRecipeByOID(recipeOid)

	if recipe == nil {
		return http.StatusNotFound, nil
	}

	ingredientOid := core.ConvertHexToObjectId(mux.Vars(r)["ingredientId"])
	log.Println(ingredientOid)

	if ingredientOid == nil {
		return http.StatusBadRequest, nil
	}

	_, ingredient := s.ingredientRepository.FindIngredientByOID(ingredientOid)

	log.Println(ingredient)
	if ingredient == nil {
		return http.StatusNotFound, nil
	}

	var ingredientDetails *IngredientDetailsDTO = &IngredientDetailsDTO{}

	invalidBody := core.DecodeBody(r, ingredientDetails)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	log.Println(ingredient)

	err := validate(ingredient, ingredientDetails)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	log.Println(ingredient.Package)
	log.Println(ingredientDetails.Metric)

	envase := getIngredientPackage(ingredientDetails.Metric, ingredient.Package)

	var recipeIngredientPackage *RecipeIngredientPackage = &RecipeIngredientPackage{
		ID: envase.ID,
	}

	var recipeIngredient *RecipeIngredient = &RecipeIngredient{
		ID:       ingredient.ID,
		Package:  *recipeIngredientPackage,
		Quantity: ingredientDetails.Quantity,
	}

	_, recipeUpdated := s.recipeRepository.AddIngredientToRecipe(recipeOid, recipeIngredient)

	if recipeUpdated == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipeUpdated
}

func validate(ingredient *ingredients.IngredientDTO, ingredientDetails *IngredientDetailsDTO) error {
	if !ingredientMetricMatches(ingredientDetails.Metric, ingredient.Package) {
		log.Println("La unidad de medida no coincide")
		return errors.New("la unidad de medida no coincide")
	}

	if ingredientDetails.Quantity == 0 {
		log.Println("La cantidad del ingrediente no puede ser 0")
		return errors.New("la cantidad del ingrediente no puede ser 0")
	}
	return nil
}

func ingredientMetricMatches(metric string, packages []packages.Package) bool {
	for _, pack := range packages {
		log.Println(pack)
		log.Println(metric)
		if fmt.Sprintf("%g %s", pack.Quantity, pack.Metric) == metric {
			return true
		}
	}
	return false
}

func getIngredientPackage(metric string, packages []packages.Package) *packages.Package {
	for _, pack := range packages {
		if fmt.Sprintf("%g %s", pack.Quantity, pack.Metric) == metric {
			return &pack
		}
	}
	return nil
}
