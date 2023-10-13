package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
)

type IngredientService struct {
	IngredientDao       dao.IngredientDao
	RecipeDao           dao.RecipeDao
	RecipeIngredientDao dao.RecipeIngredientDao
}

type IngredientServiceInterface interface {
	GetAllIngredients() (int, *[]dto.IngredientDTO)
	CreateIngredient(r *http.Request) int
	UpdateIngredient(r *http.Request) int
	DeleteIngredient(r *http.Request) int
	ChangeIngredientPrice(r *http.Request) int
}

var IngredientServiceInstance *IngredientService

func (s *IngredientService) GetAllIngredients() (int, *[]dto.IngredientDTO) {
	ingredients := s.IngredientDao.GetAllIngredients()

	return http.StatusOK, ingredients, nil
}

func (s *IngredientService) CreateIngredient(r *http.Request) int {
	ingredientDto := &dto.IngredientNameDTO{}

	err := core.DecodeBody(c, ingredientDto)

	if invalidBody {
		return http.StatusBadRequest
	}

	ingredientEntity := &models.Ingredient{
		Name:     ingredientDto.Name,
		Packages: []models.IngredientPackage{},
	}

	s.IngredientDao.CreateIngredient(ingredientEntity)

	return http.StatusCreated
}

func (s *IngredientService) UpdateIngredient(r *http.Request) int {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest
	}

	ingredient := &dto.IngredientNameDTO{}

	invalidBody := core.DecodeBody(r, ingredient)

	if invalidBody {
		return http.StatusBadRequest
	}

	err := s.IngredientDao.UpdateIngredient(oid, ingredient)

	if err != nil {
		return http.StatusInternalServerError
	}

	ingredientUpdated := s.IngredientDao.FindIngredientByOID(oid)

	if ingredientUpdated == nil {
		return http.StatusNotFound
	}

	return http.StatusOK
}

func (s *IngredientService) DeleteIngredient(r *http.Request) int {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest
	}

	err := s.IngredientDao.DeleteIngredient(oid)

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func (s *IngredientService) ChangeIngredientPrice(r *http.Request) int {
	ingredientPackageId := mux.Vars(r)["id"]
	ingredientPackageOid := core.ConvertHexToObjectId(ingredientPackageId)

	if ingredientPackageOid == nil {
		return http.StatusBadRequest
	}

	ingredientPackagePrice := &dto.IngredientPackagePriceDTO{}

	invalidBody := core.DecodeBody(r, ingredientPackagePrice)

	if invalidBody {
		return http.StatusBadRequest
	}

	err := s.IngredientDao.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)

	if err != nil {
		return http.StatusInternalServerError
	}

	err = s.IngredientDao.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)

	if ingredientUpdated == nil {
		return http.StatusInternalServerError
	}

	ingredientUpdated, err := s.IngredientDao.FindIngredientByPackageId(ingredientPackageOid)

	if len(recipes) == 0 {
		return http.StatusOK
	}

	err = s.RecipeIngredientDao.UpdateIngredientPackagePrice(ingredientPackageOid, ingredientPackagePrice.Price)

	if err != nil {
		return http.StatusInternalServerError
	}

	for i := 0; i < len(recipes); i++ {
		var recipePrice float64 = 0
		for j := 0; j < len(recipes[i].Ingredients); j++ {
			if recipes[i].Ingredients[j].ID == ingredientUpdated.ID {
				ingredientPackage := packagesById[recipes[i].Ingredients[j].Package.ID]

				recipes[i].Ingredients[j].Package = ingredientPackage
				ingredientQuantityPercent := recipes[i].Ingredients[j].Quantity / recipes[i].Ingredients[j].Package.Quantity
				recipes[i].Ingredients[j].Price = ingredientQuantityPercent * recipes[i].Ingredients[j].Package.Price
			}
			recipePrice += recipes[i].Ingredients[j].Price
		}
		recipes[i].Price = recipePrice * 3

		err := s.RecipeIngredientDao.UpdateIngredientsPrice(ingredientPackageOid, recipes[i])

		if err != nil {
			log.Println(err.Error())
			return http.StatusInternalServerError, nil, err
		}
	}

	return http.StatusOK
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
