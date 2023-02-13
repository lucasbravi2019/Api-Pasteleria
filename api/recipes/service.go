package recipes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type recipeService struct {
	recipeRepository RecipeRepository
}

type RecipeService interface {
	GetAllRecipes() (int, *[]RecipeDTO)
	GetRecipe(r *http.Request) (int, *RecipeDTO)
	CreateRecipe(r *http.Request) (int, *RecipeDTO)
	UpdateRecipeName(r *http.Request) (int, *RecipeDTO)
	DeleteRecipe(r *http.Request) (int, *primitive.ObjectID)
}

var recipeServiceInstance *recipeService

func (s *recipeService) GetAllRecipes() (int, *[]RecipeDTO) {
	recipes := s.recipeRepository.FindAllRecipes()
	return http.StatusOK, recipes
}

func (s *recipeService) GetRecipe(r *http.Request) (int, *RecipeDTO) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	recipe := s.recipeRepository.FindRecipeByOID(oid)

	if recipe == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, recipe
}

func (s *recipeService) CreateRecipe(r *http.Request) (int, *RecipeDTO) {
	var recipeName *RecipeNameDTO = &RecipeNameDTO{}

	invalidBody := core.DecodeBody(r, recipeName)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	oid := s.recipeRepository.CreateRecipe(recipeName)

	if oid == nil {
		return http.StatusInternalServerError, nil
	}

	recipe := s.recipeRepository.FindRecipeByOID(oid)

	if recipe == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusCreated, recipe
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

	err := s.recipeRepository.UpdateRecipeName(oid, recipe)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	recipeUpdated := s.recipeRepository.FindRecipeByOID(oid)

	if recipeUpdated == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipeUpdated
}

func (s *recipeService) DeleteRecipe(r *http.Request) (int, *primitive.ObjectID) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	err := s.recipeRepository.DeleteRecipe(oid)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, oid
}
