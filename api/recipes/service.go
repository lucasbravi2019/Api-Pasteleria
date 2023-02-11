package recipes

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/ingredients"
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
	ingredientOid := core.ConvertHexToObjectId(mux.Vars(r)["ingredientId"])

	if recipeOid == nil || ingredientOid == nil {
		return http.StatusBadRequest, nil
	}

	_, recipe := s.recipeRepository.FindRecipeByOID(recipeOid)

	if recipe == nil {
		return http.StatusNotFound, nil
	}

	_, ingredientDTO := s.ingredientRepository.FindIngredientByOID(ingredientOid)

	if ingredientDTO == nil {
		return http.StatusNotFound, nil
	}

	var ingredientDetails *IngredientDetailsDTO = &IngredientDetailsDTO{}

	invalidBody := core.DecodeBody(r, ingredientDetails)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	err := validate(ingredientDTO, ingredientDetails)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	envase := getIngredientPackage(ingredientDetails.Metric, ingredientDTO.Packages)

	var recipeIngredient *RecipeIngredient = &RecipeIngredient{
		ID:       primitive.NewObjectID(),
		Quantity: ingredientDetails.Quantity,
		Name:     ingredientDTO.Name,
		Package: RecipeIngredientPackage{
			ID:       envase.ID,
			Metric:   envase.Metric,
			Quantity: envase.Quantity,
			Price:    envase.Price,
		},
		Price: float64(ingredientDetails.Quantity) / envase.Quantity * envase.Price,
	}

	_, recipeUpdated := s.recipeRepository.AddIngredientToRecipe(recipeOid, recipeIngredient)

	if recipeUpdated == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, nil
}

func validate(ingredient *ingredients.IngredientDTO, ingredientDetails *IngredientDetailsDTO) error {
	if !ingredientMetricMatches(ingredientDetails.Metric, ingredient.Packages) {
		log.Println("La unidad de medida no coincide")
		return errors.New("la unidad de medida no coincide")
	}

	if ingredientDetails.Quantity == 0 {
		log.Println("La cantidad del ingrediente no puede ser 0")
		return errors.New("la cantidad del ingrediente no puede ser 0")
	}
	return nil
}

func ingredientMetricMatches(metric string, packages []ingredients.PackageDTO) bool {
	for _, pack := range packages {
		if fmt.Sprintf("%g %s", pack.Quantity, pack.Metric) == metric {
			return true
		}
	}
	return false
}

func getIngredientPackage(metric string, packages []ingredients.PackageDTO) *ingredients.PackageDTO {
	for _, pack := range packages {
		if fmt.Sprintf("%g %s", pack.Quantity, pack.Metric) == metric {
			return &pack
		}
	}
	return nil
}
