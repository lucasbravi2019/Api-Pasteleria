package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientPackageService struct {
	IngredientPackageDao dao.IngredientPackageDao
}

type IngredientPackageServiceInterface interface {
	FindAllIngredientPackages(ctx *gin.Context) (int, interface{}, error)
}

var IngredientPackageServiceInstance *IngredientPackageService

func (s *IngredientPackageService) FindAllIngredientPackages(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	ingredientId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	ingredientPackages, err := s.IngredientPackageDao.FindAllIngredientPackages(ingredientId)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, ingredientPackages, nil
}
