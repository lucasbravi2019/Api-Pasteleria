package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type PackageService struct {
	PackageDao    *dao.PackageDao
	PackageMapper *mapper.PackageMapper
	RecipeService *RecipeService
}

var PackageServiceInstance *PackageService

func (s *PackageService) GetPackages() (int, *[]dto.Package, error) {
	packages, err := s.PackageDao.GetPackages()

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, packages, nil
}

func (s *PackageService) CreatePackage(ctx *gin.Context) (int, interface{}, error) {
	var body dto.Package

	err := pkg.DecodeBody(ctx, &body)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.PackageDao.CreatePackage(&body)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (s *PackageService) UpdatePackage(ctx *gin.Context) (int, interface{}, error) {
	var body dto.Package

	err := pkg.DecodeBody(ctx, &body)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.PackageDao.UpdatePackage(&body)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	_, recipes, err := s.RecipeService.GetAllRecipes()

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	for _, recipe := range *recipes {
		err := s.RecipeService.UpdateRecipePrice(&recipe)

		if pkg.HasError(err) {
			return http.StatusInternalServerError, nil, err
		}
	}

	return http.StatusOK, nil, nil
}

func (s *PackageService) DeletePackage(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	packageId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.PackageDao.DeletePackage(&packageId)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	_, recipes, err := s.RecipeService.GetAllRecipes()

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	for _, recipe := range *recipes {
		err := s.RecipeService.UpdateRecipePrice(&recipe)

		if pkg.HasError(err) {
			return http.StatusInternalServerError, nil, err
		}
	}

	return http.StatusOK, nil, nil
}
