package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
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
	ingredients := s.IngredientDao.GetAllIngredients()

	return http.StatusOK, ingredients, nil
}

func (s *IngredientService) CreateIngredient(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusCreated, nil, nil
}

func (s *IngredientService) UpdateIngredient(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}

func (s *IngredientService) DeleteIngredient(ctx *gin.Context) (int, interface{}, error) {

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
