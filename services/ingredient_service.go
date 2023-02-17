package services

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientService struct {
	IngredientDao       dao.IngredientDao
	RecipeDao           dao.RecipeDao
	RecipeIngredientDao dao.RecipeIngredientDao
}

type IngredientServiceInterface interface {
	GetAllIngredients() (int, []dto.IngredientDTO)
	CreateIngredient(r *http.Request) (int, *dto.IngredientDTO)
	UpdateIngredient(r *http.Request) (int, *dto.IngredientDTO)
	DeleteIngredient(r *http.Request) (int, *primitive.ObjectID)
	ChangeIngredientPrice(r *http.Request) (int, *dto.IngredientDTO)
}

var IngredientServiceInstance *IngredientService

func (s *IngredientService) GetAllIngredients() (int, []dto.IngredientDTO) {
	ingredients := s.IngredientDao.GetAllIngredients()

	return http.StatusOK, ingredients
}

func (s *IngredientService) CreateIngredient(r *http.Request) (int, *dto.IngredientDTO) {
	ingredientDto := &dto.IngredientNameDTO{}

	invalidBody := core.DecodeBody(r, ingredientDto)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	ingredientEntity := &models.Ingredient{
		Name:     ingredientDto.Name,
		Packages: []models.IngredientPackage{},
	}

	ingredientCreatedId := s.IngredientDao.CreateIngredient(ingredientEntity)

	if ingredientCreatedId == nil {
		return http.StatusInternalServerError, nil
	}

	ingredientCreated := s.IngredientDao.FindIngredientByOID(ingredientCreatedId)

	if ingredientCreated == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusCreated, ingredientCreated
}

func (s *IngredientService) UpdateIngredient(r *http.Request) (int, *dto.IngredientDTO) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	ingredient := &dto.IngredientNameDTO{}

	invalidBody := core.DecodeBody(r, ingredient)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	err := s.IngredientDao.UpdateIngredient(oid, ingredient)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	ingredientUpdated := s.IngredientDao.FindIngredientByOID(oid)

	if ingredientUpdated == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, ingredientUpdated
}

func (s *IngredientService) DeleteIngredient(r *http.Request) (int, *primitive.ObjectID) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	err := s.IngredientDao.DeleteIngredient(oid)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, oid
}

func (s *IngredientService) ChangeIngredientPrice(r *http.Request) (int, *dto.IngredientDTO) {
	ingredientPackageId := mux.Vars(r)["id"]
	ingredientPackageOid := core.ConvertHexToObjectId(ingredientPackageId)

	if ingredientPackageOid == nil {
		return http.StatusBadRequest, nil
	}

	ingredientPackagePrice := &dto.IngredientPackagePriceDTO{}

	invalidBody := core.DecodeBody(r, ingredientPackagePrice)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	err := s.IngredientDao.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	ingredientUpdated := s.IngredientDao.FindIngredientByPackageId(ingredientPackageOid)

	if ingredientUpdated == nil {
		return http.StatusInternalServerError, nil
	}

	recipes := s.RecipeDao.FindRecipesByPackageId(ingredientPackageOid)

	if len(recipes) == 0 {
		return http.StatusOK, ingredientUpdated
	}

	err = s.RecipeIngredientDao.UpdateIngredientPackagePrice(ingredientPackageOid, ingredientPackagePrice.Price)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	for i := 0; i < len(recipes); i++ {
		var recipePrice float64 = 0
		for j := 0; j < len(recipes[i].Ingredients); j++ {
			recipes[i].Ingredients[j].Price = recipes[i].Ingredients[j].Quantity / recipes[i].Ingredients[j].Package.Quantity * recipes[i].Ingredients[j].Package.Price
			recipePrice += recipes[i].Ingredients[j].Price
		}
		recipes[i].Price = recipePrice * 3

		err := s.RecipeIngredientDao.UpdateIngredientsPrice(ingredientPackageOid, recipes[i])

		if err != nil {
			log.Println(err.Error())
		}
	}

	return http.StatusOK, ingredientUpdated
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
