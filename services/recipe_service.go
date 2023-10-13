package services

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
)

type RecipeService struct {
	RecipeDao dao.RecipeDao
}

type RecipeServiceInterface interface {
	GetAllRecipes() (int, *[]dto.RecipeDTO)
	GetRecipe(r *http.Request) (int, *dto.RecipeDTO)
	CreateRecipe(r *http.Request) error
	UpdateRecipeName(r *http.Request) int
	DeleteRecipe(r *http.Request) int
}

var RecipeServiceInstance *RecipeService

func (s *RecipeService) GetAllRecipes() (int, *[]dto.RecipeDTO) {
	recipes := s.RecipeDao.FindAllRecipes()
	return http.StatusOK, recipes
}

func (s *RecipeService) GetRecipe(r *http.Request) (int, *dto.RecipeDTO) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	recipe := s.RecipeDao.FindRecipeByOID(oid)

	if recipe == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, recipe
}

func (s *RecipeService) CreateRecipe(r *http.Request) error {
	recipeName := &dto.RecipeNameDTO{}

	invalidBody := core.DecodeBody(r, recipeName)

	if invalidBody {
		return errors.New("body request not valid")
	}

	err := s.RecipeDao.CreateRecipe(recipeName)

	return err
}

func (s *RecipeService) UpdateRecipeName(r *http.Request) int {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest
	}

	recipe := &dto.RecipeNameDTO{}

	invalidBody := core.DecodeBody(r, recipe)

	if invalidBody {
		return http.StatusBadRequest
	}

	err := s.RecipeDao.UpdateRecipeName(oid, recipe)

	if err != nil {
		return http.StatusInternalServerError
	}

	recipeUpdated := s.RecipeDao.FindRecipeByOID(oid)

	if recipeUpdated == nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func (s *RecipeService) DeleteRecipe(r *http.Request) int {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest
	}

	err := s.RecipeDao.DeleteRecipe(oid)

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
