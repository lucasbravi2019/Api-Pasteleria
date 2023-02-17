package services

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientPackageService struct {
	PackageDao           dao.PackageDao
	IngredientPackageDao dao.IngredientPackageDao
}

type IngredientPackageServiceInterface interface {
	AddPackageToIngredient(r *http.Request) int
	RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID)
}

var IngredientPackageServiceInstance *IngredientPackageService

func (s *IngredientPackageService) AddPackageToIngredient(r *http.Request) int {
	ingredientOid := mux.Vars(r)["ingredientId"]
	packageOid := mux.Vars(r)["packageId"]
	ingredientId := core.ConvertHexToObjectId(ingredientOid)
	packageId := core.ConvertHexToObjectId(packageOid)

	priceDTO := &dto.IngredientPackagePriceDTO{}

	invalidBody := core.DecodeBody(r, priceDTO)

	if invalidBody {
		return http.StatusBadRequest
	}

	envase := s.PackageDao.GetPackageById(packageId)

	if envase == nil {
		return http.StatusNotFound
	}

	ingredientPackage := &models.IngredientPackage{
		ID:       envase.ID,
		Metric:   envase.Metric,
		Quantity: envase.Quantity,
		Price:    priceDTO.Price,
	}

	err := s.IngredientPackageDao.AddPackageToIngredient(ingredientId, packageId, ingredientPackage)

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func (s *IngredientPackageService) RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID) {
	packageOid := core.ConvertHexToObjectId(mux.Vars(r)["packageId"])

	ingredientPackageDto := &dto.IngredientPackageDTO{
		PackageOid: *packageOid,
	}

	err := s.IngredientPackageDao.RemovePackageFromIngredients(*ingredientPackageDto)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, packageOid
}
