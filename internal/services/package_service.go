package services

import (
	"net/http"

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
	CreatePackage(r *http.Request) int
	UpdatePackage(r *http.Request) int
	DeletePackage(r *http.Request) int
}

var PackageServiceInstance *PackageService

func (s *PackageService) GetPackages() (int, *[]models.Package, error) {

	return http.StatusOK, nil, nil
}

func (s *PackageService) CreatePackage(r *http.Request) int {

	return http.StatusCreated
}

func (s *PackageService) UpdatePackage(r *http.Request) int {

	return http.StatusOK
}

func (s *PackageService) DeletePackage(r *http.Request) int {

	return http.StatusOK
}
