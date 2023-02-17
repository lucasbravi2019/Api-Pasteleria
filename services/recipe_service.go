package services

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeService struct {
	RecipeDao dao.RecipeDao
}

type RecipeServiceInterface interface {
	GetAllRecipes() (int, *[]dto.RecipeDTO)
	GetRecipe(r *http.Request) (int, *dto.RecipeDTO)
	CreateRecipe(r *http.Request) (int, *dto.RecipeDTO)
	UpdateRecipeName(r *http.Request) (int, *dto.RecipeDTO)
	DeleteRecipe(r *http.Request) (int, *primitive.ObjectID)
}

var RecipeServiceInstance *RecipeService

func (s *RecipeService) GetAllRecipes() (int, *[]dto.RecipeDTO) {
	recipes := s.RecipeDao.FindAllRecipes()
	return http.StatusOK, recipes
}

func (s *RecipeService) GetRecipe(r *http.Request) (int, *dto.RecipeDTO) {
	recipeId := &dto.RecipeIdDTO{}

	invalidBody := core.DecodeBody(r, recipeId)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	oid := core.ConvertHexToObjectId(recipeId.ID)

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	recipe := s.RecipeDao.FindRecipeByOID(oid)

	if recipe == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, recipe
}

func (s *RecipeService) CreateRecipe(r *http.Request) (int, *dto.RecipeDTO) {
	recipeName := &dto.RecipeNameDTO{}

	invalidBody := core.DecodeBody(r, recipeName)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	recipeEntity := &models.Recipe{
		Name: recipeName.Name,
	}

	oid := s.RecipeDao.CreateRecipe(recipeEntity)

	if oid == nil {
		return http.StatusInternalServerError, nil
	}

	recipeCreated := s.RecipeDao.FindRecipeByOID(oid)

	if recipeCreated == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusCreated, recipeCreated
}

func (s *RecipeService) UpdateRecipeName(r *http.Request) (int, *dto.RecipeDTO) {
	recipeName := &dto.RecipeNameDTO{}

	invalidBody := core.DecodeBody(r, recipeName)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	oid := core.ConvertHexToObjectId(recipeName.ID)

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	recipe := &models.Recipe{
		Name: recipeName.Name,
	}

	err := s.RecipeDao.UpdateRecipeName(oid, recipe)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	recipeUpdated := s.RecipeDao.FindRecipeByOID(oid)

	if recipeUpdated == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipeUpdated
}

func (s *RecipeService) DeleteRecipe(r *http.Request) (int, *primitive.ObjectID) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	err := s.RecipeDao.DeleteRecipe(oid)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, oid
}
