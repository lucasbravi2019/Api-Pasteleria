package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
)

type RecipeIngredientService struct {
	RecipeDao           dao.RecipeDao
	IngredientDao       dao.IngredientDao
	RecipeIngredientDao dao.RecipeIngredientDao
	RecipeMapper        mapper.RecipeMapper
}

type RecipeIngredientServiceInterface interface {
	AddIngredientToRecipe(r *http.Request) (int, error)
	RemoveIngredientFromRecipe(r *http.Request) (int, *dto.RecipeDTO, error)
}

var RecipeIngredientServiceInstance *RecipeIngredientService

func (s *RecipeIngredientService) AddIngredientToRecipe(r *http.Request) (int, error) {

	return http.StatusOK, nil
}

func (s *RecipeIngredientService) RemoveIngredientFromRecipe(c *gin.Context) (int, *dto.RecipeDTO, error) {

	return http.StatusOK, nil, nil
}
