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

type PackageService struct {
	PackageDao           dao.PackageDao
	IngredientPackageDao dao.IngredientPackageDao
	RecipeDao            dao.RecipeDao
	RecipeIngredientDao  dao.RecipeIngredientDao
}

type PackageServiceInterface interface {
	GetPackages() (int, *[]models.Package, error)
	CreatePackage(c *gin.Context) (int, *models.Package, error)
	UpdatePackage(c *gin.Context) (int, *models.Package, error)
	DeletePackage(c *gin.Context) (int, *primitive.ObjectID, error)
}

var PackageServiceInstance *PackageService

func (s *PackageService) GetPackages() (int, *[]models.Package, error) {
	packages, err := s.PackageDao.GetPackages()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, packages, nil
}

func (s *PackageService) CreatePackage(c *gin.Context) (int, *models.Package, error) {
	packageRequest := &models.Package{}

	err := core.DecodeBody(c, packageRequest)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	id, err := s.PackageDao.CreatePackage(packageRequest)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	envase, err := s.PackageDao.GetPackageById(id)

	if err != nil {
		return http.StatusNotFound, nil, err
	}

	return http.StatusCreated, envase, nil
}

func (s *PackageService) UpdatePackage(c *gin.Context) (int, *models.Package, error) {
	oid, err := core.ConvertUrlVarToObjectId("packageId", c)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	packageRequest := &models.Package{}

	err = core.DecodeBody(c, packageRequest)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.PackageDao.UpdatePackage(oid, packageRequest)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	envase, err := s.PackageDao.GetPackageById(oid)

	if err != nil {
		return http.StatusNotFound, nil, err
	}

	return http.StatusOK, envase, nil
}

func (s *PackageService) DeletePackage(c *gin.Context) (int, *primitive.ObjectID, error) {
	oid, err := core.ConvertUrlVarToObjectId("packageId", c)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.PackageDao.DeletePackage(oid)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	ingredientPackage := &dto.IngredientPackageDTO{
		PackageOid: *oid,
	}

	err = s.IngredientPackageDao.RemovePackageFromIngredients(*ingredientPackage)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	err = s.RecipeIngredientDao.RemoveIngredientByPackageId(oid)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	err = s.RecipeDao.UpdateRecipesPrice()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, oid, nil
}
