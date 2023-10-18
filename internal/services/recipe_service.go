package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type RecipeService struct {
	RecipeDao dao.RecipeDao
}

type RecipeServiceInterface interface {
	GetAllRecipes() (int, *[]dto.RecipeDTO, error)
	GetRecipe(ctx *gin.Context) (int, *dto.RecipeDTO, error)
	CreateRecipe(ctx *gin.Context) (int, interface{}, error)
	UpdateRecipeName(ctx *gin.Context) (int, interface{}, error)
	DeleteRecipe(ctx *gin.Context) (int, interface{}, error)
}

var RecipeServiceInstance *RecipeService

func (s *RecipeService) GetAllRecipes() (int, *[]dto.RecipeDTO, error) {
	recipes, err := s.RecipeDao.FindAllRecipes()

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	dtos := mapper.ToRecipeDTOList(recipes)

	return http.StatusOK, dtos, nil
}

func (s *RecipeService) GetRecipe(ctx *gin.Context) (int, *dto.RecipeDTO, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}
	recipeId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	recipe, err := s.RecipeDao.FindRecipeById(recipeId)

	if pkg.HasError(err) {
		return http.StatusNotFound, nil, err
	}

	if recipe == nil {
		return http.StatusNotFound, nil, fmt.Errorf("recipe not found for id: %d", recipeId)
	}

	dto := mapper.ToRecipeDTO(*recipe)

	return http.StatusOK, dto, nil
}

func (s *RecipeService) CreateRecipe(ctx *gin.Context) (int, interface{}, error) {
	var recipeName dto.RecipeNameDTO

	err := pkg.DecodeBody(ctx, &recipeName)

	if pkg.HasError(err) {
		log.Println(err)
		return http.StatusBadRequest, nil, err
	}

	err = s.RecipeDao.CreateRecipe(&recipeName)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (s *RecipeService) UpdateRecipeName(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	recipeId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	var recipeName dto.RecipeNameDTO

	err = pkg.DecodeBody(ctx, &recipeName)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	err = s.RecipeDao.UpdateRecipeName(&recipeId, &recipeName)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}

func (s *RecipeService) DeleteRecipe(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	recipeId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	err = s.RecipeDao.DeleteRecipe(&recipeId)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}
