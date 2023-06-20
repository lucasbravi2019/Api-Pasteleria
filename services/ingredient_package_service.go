package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
	AddPackageToIngredient(c *gin.Context) (int, error)
	RemovePackageFromIngredients(c *gin.Context) (int, *primitive.ObjectID, error)
}

var IngredientPackageServiceInstance *IngredientPackageService

func (s *IngredientPackageService) AddPackageToIngredient(c *gin.Context) (int, error) {
	ingredientId, err := core.ConvertUrlVarToObjectId("ingredientId", c)

	if err != nil {
		return http.StatusBadRequest, err
	}

	packageId, err := core.ConvertUrlVarToObjectId("packageId", c)

	if err != nil {
		return http.StatusBadRequest, err
	}

	priceDTO := &dto.IngredientPackagePriceDTO{}

	err = core.DecodeBody(c, priceDTO)

	if err != nil {
		return http.StatusBadRequest, err
	}

	envase, err := s.PackageDao.GetPackageById(packageId)

	if err != nil {
		return http.StatusNotFound, err
	}

	ingredientPackage := &models.IngredientPackage{
		ID:       envase.ID,
		Metric:   envase.Metric,
		Quantity: envase.Quantity,
		Price:    priceDTO.Price,
	}

	err = s.IngredientPackageDao.AddPackageToIngredient(ingredientId, packageId, ingredientPackage)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *IngredientPackageService) RemovePackageFromIngredients(c *gin.Context) (int, *primitive.ObjectID, error) {
	packageOid, err := core.ConvertUrlVarToObjectId("packageId", c)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredientPackageDto := &dto.IngredientPackageDTO{
		PackageOid: *packageOid,
	}

	err = s.IngredientPackageDao.RemovePackageFromIngredients(*ingredientPackageDto)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, packageOid, nil
}
