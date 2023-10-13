package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
)

type PackageService struct {
	PackageDao           dao.PackageDao
	IngredientPackageDao dao.IngredientPackageDao
	RecipeDao            dao.RecipeDao
	RecipeIngredientDao  dao.RecipeIngredientDao
}

type PackageServiceInterface interface {
	GetPackages() (int, *[]models.Package)
	CreatePackage(r *http.Request) int
	UpdatePackage(r *http.Request) int
	DeletePackage(r *http.Request) int
}

var PackageServiceInstance *PackageService

func (s *PackageService) GetPackages() (int, *[]models.Package, error) {
	packages, err := s.PackageDao.GetPackages()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, packages, nil
}

func (s *PackageService) CreatePackage(r *http.Request) int {
	packageRequest := &models.Package{}

	err := core.DecodeBody(c, packageRequest)

	if invalidBody {
		return http.StatusBadRequest
	}

	s.PackageDao.CreatePackage(packageRequest)

	return http.StatusCreated
}

func (s *PackageService) UpdatePackage(r *http.Request) int {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest
	}

	packageRequest := &models.Package{}

	invalidBody := core.DecodeBody(r, packageRequest)

	if invalidBody {
		return http.StatusBadRequest
	}

	err := s.PackageDao.UpdatePackage(oid, packageRequest)

	if err != nil {
		return http.StatusInternalServerError
	}

	err = s.PackageDao.UpdatePackage(oid, packageRequest)

	if envase == nil {
		return http.StatusNotFound
	}

	return http.StatusOK
}

func (s *PackageService) DeletePackage(r *http.Request) int {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest
	}

	err := s.PackageDao.DeletePackage(oid)

	if err != nil {
		return http.StatusInternalServerError
	}

	ingredientPackage := &dto.IngredientPackageDTO{
		PackageOid: *oid,
	}

	err = s.IngredientPackageDao.RemovePackageFromIngredients(*ingredientPackage)

	if err != nil {
		return http.StatusInternalServerError
	}

	err = s.RecipeIngredientDao.RemoveIngredientByPackageId(oid)

	if err != nil {
		return http.StatusInternalServerError
	}

	err = s.RecipeDao.UpdateRecipesPrice()

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
