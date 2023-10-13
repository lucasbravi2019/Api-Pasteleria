package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
)

type IngredientPackageService struct {
	PackageDao           dao.PackageDao
	IngredientPackageDao dao.IngredientPackageDao
}

type IngredientPackageServiceInterface interface {
	AddPackageToIngredient(ctx *gin.Context) (int, interface{}, error)
	RemovePackageFromIngredients(ctx *gin.Context) (int, interface{}, error)
}

var IngredientPackageServiceInstance *IngredientPackageService

func (s *IngredientPackageService) AddPackageToIngredient(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}

func (s *IngredientPackageService) RemovePackageFromIngredients(ctx *gin.Context) (int, interface{}, error) {

	return http.StatusOK, nil, nil
}
