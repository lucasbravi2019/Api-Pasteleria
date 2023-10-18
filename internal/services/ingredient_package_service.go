package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientPackageService struct {
	PackageDao           dao.PackageDao
	IngredientPackageDao dao.IngredientPackageDao
}

type IngredientPackageServiceInterface interface {
	AddPackageToIngredient(ctx *gin.Context) (int, interface{}, error)
	RemovePackageFromIngredients(ctx *gin.Context) (int, interface{}, error)
	FindAllIngredientPackages(ctx *gin.Context) (int, interface{}, error)
}

var IngredientPackageServiceInstance *IngredientPackageService

func (s *IngredientPackageService) AddPackageToIngredient(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}

func (s *IngredientPackageService) RemovePackageFromIngredients(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}

func (s *IngredientPackageService) FindAllIngredientPackages(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredientId, err := util.ToLong(id)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredientPackages, err := s.IngredientPackageDao.FindAllIngredientPackages(ingredientId)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, ingredientPackages, nil
}
