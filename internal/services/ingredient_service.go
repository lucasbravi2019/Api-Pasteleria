package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientService struct {
	IngredientDao dao.IngredientDao
}

var IngredientServiceInstance *IngredientService

func (s *IngredientService) GetAllIngredients() (int, *[]dto.IngredientResponse, error) {
	ingredients, err := s.IngredientDao.GetAllIngredients()

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, ingredients, nil
}

func (s *IngredientService) CreateIngredient(ctx *gin.Context) (int, interface{}, error) {
	var ingredient dto.IngredientRequest

	err := pkg.DecodeBody(ctx, &ingredient)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.CreateIngredientName(ingredient.Name)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	ingredientId, err := s.IngredientDao.FindIngredientIdByName(ingredient.Name)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	err = s.IngredientDao.AddIngredientPackage(ingredientId, ingredient.Packages)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (s *IngredientService) UpdateIngredient(ctx *gin.Context) (int, interface{}, error) {
	var ingredient dto.IngredientRequest

	err := pkg.DecodeBody(ctx, &ingredient)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.UpdateIngredientName(&ingredient)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	err = s.IngredientDao.RemoveIngredientPackages(ingredient.Id)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	err = s.IngredientDao.AddIngredientPackage(ingredient.Id, ingredient.Packages)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}

func (s *IngredientService) DeleteIngredient(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	ingredientId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.DeleteIngredient(&ingredientId)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}
