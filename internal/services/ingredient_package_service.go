package services

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientPackageService struct {
	PackageDao           dao.PackageDao
	IngredientPackageDao dao.IngredientPackageDao
}

type IngredientPackageServiceInterface interface {
	AddPackageToIngredient(r *http.Request) (int, error)
	RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID, error)
}

var IngredientPackageServiceInstance *IngredientPackageService

func (s *IngredientPackageService) AddPackageToIngredient(r *http.Request) (int, error) {

	return http.StatusOK, nil
}

func (s *IngredientPackageService) RemovePackageFromIngredients(ctx *http.Request) (int, *primitive.ObjectID, error) {

	return http.StatusOK, nil, nil
}
