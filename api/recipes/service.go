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
