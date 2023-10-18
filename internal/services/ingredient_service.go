package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientService struct {
	IngredientDao       dao.IngredientDao
	RecipeDao           dao.RecipeDao
	RecipeIngredientDao dao.RecipeIngredientDao
}

type IngredientServiceInterface interface {
	GetAllIngredients() (int, *[]dto.IngredientDTO, error)
	CreateIngredient(ctx *gin.Context) (int, interface{}, error)
	UpdateIngredient(ctx *gin.Context) (int, interface{}, error)
	DeleteIngredient(ctx *gin.Context) (int, interface{}, error)
	ChangeIngredientPrice(ctx *gin.Context) (int, interface{}, error)
}

var IngredientServiceInstance *IngredientService

func (s *IngredientService) GetAllIngredients() (int, *[]dto.IngredientDTO, error) {
	ingredients, err := s.IngredientDao.GetAllIngredients()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, ingredients, nil
}

func (s *IngredientService) CreateIngredient(ctx *gin.Context) (int, interface{}, error) {
	var ingredient dto.IngredientNameDTO

	err := pkg.DecodeBody(ctx, &ingredient)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.CreateIngredient(&ingredient)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (s *IngredientService) UpdateIngredient(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredientId, err := util.ToLong(id)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	var ingredientName dto.IngredientNameDTO

	err = pkg.DecodeBody(ctx, &ingredientName)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.UpdateIngredient(ingredientId, &ingredientName)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}

func (s *IngredientService) DeleteIngredient(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredientId, err := util.ToLong(id)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.DeleteIngredient(&ingredientId)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}

func (s *IngredientService) ChangeIngredientPrice(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}

func validate(ingredient *dto.IngredientDTO, ingredientDetails *dto.IngredientDetailsDTO) error {
	if !ingredientMetricMatches(ingredientDetails.Metric, ingredient.Packages) {
		log.Println("La unidad de medida no coincide")
		return errors.New("la unidad de medida no coincide")
	}

	if ingredientDetails.Quantity == 0 {
		log.Println("La cantidad del ingrediente no puede ser 0")
		return errors.New("la cantidad del ingrediente no puede ser 0")
	}
	return nil
}

func ingredientMetricMatches(metric string, packages []dto.PackageDTO) bool {
	for _, pack := range packages {
		if fmt.Sprintf("%g %s", pack.Quantity, pack.Metric) == metric {
			return true
		}
	}
	return false
}

func getIngredientPackage(metric string, packages []dto.PackageDTO) *dto.PackageDTO {
	for _, pack := range packages {
		if fmt.Sprintf("%g %s", pack.Quantity, pack.Metric) == metric {
			return &pack
		}
	}
	return nil
}
