package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/models"
)

type PackageService struct {
	PackageDao           dao.PackageDao
	IngredientPackageDao dao.IngredientPackageDao
	RecipeDao            dao.RecipeDao
	RecipeIngredientDao  dao.RecipeIngredientDao
}

type PackageServiceInterface interface {
	GetPackages() (int, *[]models.Package, error)
	CreatePackage(ctx *gin.Context) (int, interface{}, error)
	UpdatePackage(ctx *gin.Context) (int, interface{}, error)
	DeletePackage(ctx *gin.Context) (int, interface{}, error)
}

var PackageServiceInstance *PackageService

func (s *PackageService) GetPackages() (int, *[]models.Package, error) {

	return http.StatusOK, nil, nil
}

func (s *PackageService) CreatePackage(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusCreated, nil, nil
}

func (s *PackageService) UpdatePackage(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}

func (s *PackageService) DeletePackage(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}
