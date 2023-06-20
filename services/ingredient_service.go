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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientService struct {
	IngredientDao       dao.IngredientDao
	RecipeDao           dao.RecipeDao
	RecipeIngredientDao dao.RecipeIngredientDao
}

type IngredientServiceInterface interface {
	GetAllIngredients() (int, []dto.IngredientDTO, error)
	CreateIngredient(c *gin.Context) (int, *dto.IngredientDTO, error)
	UpdateIngredient(c *gin.Context) (int, *dto.IngredientDTO, error)
	DeleteIngredient(c *gin.Context) (int, *primitive.ObjectID, error)
	ChangeIngredientPrice(c *gin.Context) (int, *dto.IngredientDTO, error)
}

var IngredientServiceInstance *IngredientService

func (s *IngredientService) GetAllIngredients() (int, []dto.IngredientDTO, error) {
	ingredients := s.IngredientDao.GetAllIngredients()

	return http.StatusOK, ingredients, nil
}

func (s *IngredientService) CreateIngredient(c *gin.Context) (int, *dto.IngredientDTO, error) {
	ingredientDto := &dto.IngredientNameDTO{}

	err := core.DecodeBody(c, ingredientDto)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredientEntity := &models.Ingredient{
		Name:     ingredientDto.Name,
		Packages: []models.IngredientPackage{},
	}

	ingredientCreatedId, err := s.IngredientDao.CreateIngredient(ingredientEntity)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	ingredientCreated, err := s.IngredientDao.FindIngredientByOID(ingredientCreatedId)

	if err != nil {
		return http.StatusNotFound, nil, err
	}

	return http.StatusCreated, ingredientCreated, nil
}

func (s *IngredientService) UpdateIngredient(c *gin.Context) (int, *dto.IngredientDTO, error) {
	oid, err := core.ConvertUrlVarToObjectId("ingredientId", c)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredient := &dto.IngredientNameDTO{}

	err = core.DecodeBody(c, ingredient)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.UpdateIngredient(oid, ingredient)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	ingredientUpdated, err := s.IngredientDao.FindIngredientByOID(oid)

	if err != nil {
		return http.StatusNotFound, nil, err
	}

	recipes, err := s.RecipeDao.GetRecipesByIngredientId(oid)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	for i := 0; i < len(*recipes); i++ {
		for j := 0; j < len((*recipes)[i].Ingredients); j++ {
			ingredient := &(*recipes)[i].Ingredients[j]
			if ingredient.ID == ingredientUpdated.ID {
				ingredient.Name = ingredientUpdated.Name
			}
		}
	}

	err = s.RecipeDao.UpdateRecipes(*recipes)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, ingredientUpdated, nil
}

func (s *IngredientService) DeleteIngredient(c *gin.Context) (int, *primitive.ObjectID, error) {
	oid, err := core.ConvertUrlVarToObjectId("ingredientId", c)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.DeleteIngredient(oid)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, oid, nil
}

func (s *IngredientService) ChangeIngredientPrice(c *gin.Context) (int, *dto.IngredientDTO, error) {
	ingredientPackageOid, err := core.ConvertUrlVarToObjectId("ingredientId", c)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredientPackagePrice := &dto.IngredientPackagePriceDTO{}

	err = core.DecodeBody(c, ingredientPackagePrice)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	ingredientUpdated, err := s.IngredientDao.FindIngredientByPackageId(ingredientPackageOid)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	recipes, err := s.RecipeDao.FindRecipesByPackageId(ingredientPackageOid)

	if err != nil {
		return http.StatusNotFound, nil, err
	}

	err = s.RecipeIngredientDao.UpdateIngredientPackagePrice(ingredientPackageOid, ingredientPackagePrice.Price)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	packagesById := make(map[primitive.ObjectID]dto.PackageDTO)

	for i := 0; i < len(ingredientUpdated.Packages); i++ {
		packagesById[ingredientUpdated.Packages[i].ID] = ingredientUpdated.Packages[i]
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

	return http.StatusOK, ingredientUpdated, nil
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
