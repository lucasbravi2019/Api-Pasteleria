package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type RecipeIngredientService struct {
	RecipeIngredientDao dao.RecipeIngredientDao
}

type RecipeIngredientInterface interface {
	GetAllRecipeIngredients(ctx *gin.Context) (int, interface{}, error)
	UpdateRecipeIngredients(ctx *gin.Context) (int, interface{}, error)
}

var RecipeIngredientServiceInstance *RecipeIngredientService

func (s *RecipeIngredientService) GetAllRecipeIngredients(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlParams(ctx, "id")

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	recipeId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	recipes, err := s.RecipeIngredientDao.GetAllRecipeIngredients(recipeId)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, recipes, nil
}

func (s *RecipeIngredientService) UpdateRecipeIngredients(ctx *gin.Context) (int, interface{}, error) {
	var ingredients dto.RecipeIngredientIdDTO

	err := pkg.DecodeBody(ctx, &ingredients)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.RecipeIngredientDao.UpdateRecipeIngredients(ingredients)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}
