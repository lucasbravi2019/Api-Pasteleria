package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type RecipeService struct {
	RecipeDao    *dao.RecipeDao
	RecipeMapper *mapper.RecipeMapper
}

type RecipeServiceInterface interface {
	GetAllRecipes(ctx *gin.Context) (int, *[]dto.RecipeDTO, error)
	GetRecipe(ctx *gin.Context) (int, *dto.RecipeDTO, error)
	CreateRecipe(ctx *gin.Context) (int, interface{}, error)
	UpdateRecipe(ctx *gin.Context) (int, interface{}, error)
	DeleteRecipe(ctx *gin.Context) (int, interface{}, error)
}

var RecipeServiceInstance *RecipeService

func (s *RecipeService) GetAllRecipes(ctx *gin.Context) (int, *[]dto.RecipeDTO, error) {
	recipesFound, err := s.RecipeDao.FindAllRecipes()

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, recipesFound, nil
}

func (s *RecipeService) GetRecipe(ctx *gin.Context) (int, *dto.RecipeDTO, error) {
	id, _ := pkg.GetUrlVars(ctx, "id")

	recipeId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	recipe, err := s.RecipeDao.FindRecipeById(recipeId)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, recipe, nil
}

func (s *RecipeService) CreateRecipe(ctx *gin.Context) (int, interface{}, error) {
	var recipe dto.RecipeCreationDTO

	err := pkg.DecodeBody(ctx, &recipe)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.RecipeDao.CreateRecipe(&recipe)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (s *RecipeService) UpdateRecipe(ctx *gin.Context) (int, interface{}, error) {
	var recipe dto.RecipeUpdateDTO

	err := pkg.DecodeBody(ctx, &recipe)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	err = s.RecipeDao.UpdateRecipe(&recipe)

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
