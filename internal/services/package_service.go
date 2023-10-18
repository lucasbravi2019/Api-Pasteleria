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
	PackageDao           dao.PackageDao
	IngredientPackageDao dao.IngredientPackageDao
	RecipeDao            dao.RecipeDao
	RecipeIngredientDao  dao.RecipeIngredientDao
}

type PackageServiceInterface interface {
	GetPackages() (int, *[]dto.PackageDTO, error)
	CreatePackage(ctx *gin.Context) (int, interface{}, error)
	UpdatePackage(ctx *gin.Context) (int, interface{}, error)
	DeletePackage(ctx *gin.Context) (int, interface{}, error)
}

var PackageServiceInstance *PackageService

func (s *PackageService) GetPackages() (int, *[]dto.PackageDTO, error) {
	packages, err := s.PackageDao.GetPackages()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	dtos := mapper.ToPackageDTOList(packages)

	return http.StatusOK, dtos, nil
}

func (s *PackageService) CreatePackage(ctx *gin.Context) (int, interface{}, error) {
	var body dto.PackageDTO

	err := pkg.DecodeBody(ctx, &body)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.PackageDao.CreatePackage(&body)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (s *PackageService) UpdatePackage(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	packageId, err := util.ToLong(id)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	var body dto.PackageDTO

	err = pkg.DecodeBody(ctx, &body)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.PackageDao.UpdatePackage(&packageId, &body)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}

func (s *PackageService) DeletePackage(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	packageId, err := util.ToLong(id)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	err = s.PackageDao.DeletePackage(&packageId)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, nil, nil
}
