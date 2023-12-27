package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type RecipeService struct {
	RecipeDao    *dao.RecipeDao
	RecipeMapper *mapper.RecipeMapper
}

type RecipeServiceInterface interface {
	GetAllRecipes(ctx *gin.Context) (int, *[]dto.RecipeDTO, error)
	CreateRecipe(ctx *gin.Context) (int, interface{}, error)
	UpdateRecipeName(ctx *gin.Context) (int, interface{}, error)
	DeleteRecipe(ctx *gin.Context) (int, interface{}, error)
}

var RecipeServiceInstance *RecipeService

func (s *RecipeService) GetAllRecipes(ctx *gin.Context) (int, *[]dto.RecipeDTO, error) {
	id, _ := pkg.GetUrlParams(ctx, "id")

	recipes := util.NewList[models.Recipe]()
	var err error

	if pkg.STRING_EMPTY != id {
		recipeId, err := util.ToLong(id)

		if pkg.HasError(err) {
			return http.StatusBadRequest, nil, err
		}

		recipe, err := s.RecipeDao.FindRecipeById(recipeId)

		if pkg.HasError(err) {
			return http.StatusInternalServerError, nil, err
		}

		util.AddAll(&recipes, *recipe)
	} else {
		recipesFound, err := s.RecipeDao.FindAllRecipes()

		if pkg.HasError(err) {
			return http.StatusInternalServerError, nil, err
		}

		util.AddAll(&recipes, *recipesFound)
	}

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	dtos := s.RecipeMapper.ToRecipeDTOList(&recipes)

	return http.StatusOK, dtos, nil
}

func (s *RecipeService) CreateRecipe(ctx *gin.Context) (int, interface{}, error) {
	var recipeName dto.RecipeNameDTO

	err := pkg.DecodeBody(ctx, &recipeName)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.RecipeDao.CreateRecipe(&recipeName)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (s *RecipeService) UpdateRecipeName(ctx *gin.Context) (int, interface{}, error) {
	var recipeName dto.RecipeNameDTO

	err := pkg.DecodeBody(ctx, &recipeName)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	err = s.RecipeDao.UpdateRecipeName(&recipeName)

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
