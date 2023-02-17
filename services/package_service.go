package services

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PackageService struct {
	PackageDao    dao.PackageDao
	IngredientDao dao.IngredientDao
	RecipeDao     dao.RecipeDao
}

type PackageServiceInterface interface {
	GetPackages() (int, *[]models.Package)
	CreatePackage(r *http.Request) (int, *models.Package)
	UpdatePackage(r *http.Request) (int, *models.Package)
	DeletePackage(r *http.Request) (int, *primitive.ObjectID)
	AddPackageToIngredient(r *http.Request) int
	RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID)
}

var PackageServiceInstance *PackageService

func (s *PackageService) GetPackages() (int, *[]models.Package) {
	packages := s.PackageDao.GetPackages()

	if packages == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, packages
}

func (s *PackageService) CreatePackage(r *http.Request) (int, *models.Package) {
	packageRequest := &models.Package{}

	invalidBody := core.DecodeBody(r, packageRequest)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	id := s.PackageDao.CreatePackage(packageRequest)

	if id == nil {
		return http.StatusInternalServerError, nil
	}

	envase := s.PackageDao.GetPackageById(id)

	if envase == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusCreated, envase
}

func (s *PackageService) UpdatePackage(r *http.Request) (int, *models.Package) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	packageRequest := &models.Package{}

	invalidBody := core.DecodeBody(r, packageRequest)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	err := s.PackageDao.UpdatePackage(oid, packageRequest)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	envase := s.PackageDao.GetPackageById(oid)

	if envase == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, envase
}

func (s *PackageService) DeletePackage(r *http.Request) (int, *primitive.ObjectID) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	err := s.PackageDao.DeletePackage(oid)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	ingredientPackage := &dto.IngredientPackageDTO{
		PackageOid: *oid,
	}

	err = s.IngredientDao.RemovePackageFromIngredients(*ingredientPackage)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	err = s.RecipeDao.RemoveIngredientByPackageId(oid)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	err = s.RecipeDao.UpdateRecipesPrice()

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, oid
}

func (s *PackageService) AddPackageToIngredient(r *http.Request) int {
	ingredientOid := mux.Vars(r)["ingredientId"]
	packageOid := mux.Vars(r)["packageId"]
	ingredientId := core.ConvertHexToObjectId(ingredientOid)
	packageId := core.ConvertHexToObjectId(packageOid)

	priceDTO := &dto.IngredientPackagePriceDTO{}

	invalidBody := core.DecodeBody(r, priceDTO)

	if invalidBody {
		return http.StatusBadRequest
	}

	envase := s.PackageDao.GetPackageById(packageId)

	if envase == nil {
		return http.StatusNotFound
	}

	ingredientPackage := &models.IngredientPackage{
		ID:       envase.ID,
		Metric:   envase.Metric,
		Quantity: envase.Quantity,
		Price:    priceDTO.Price,
	}

	err := s.IngredientDao.AddPackageToIngredient(ingredientId, packageId, ingredientPackage)

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func (s *PackageService) RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID) {
	packageOid := core.ConvertHexToObjectId(mux.Vars(r)["packageId"])

	ingredientPackageDto := &dto.IngredientPackageDTO{
		PackageOid: *packageOid,
	}

	err := s.IngredientDao.RemovePackageFromIngredients(*ingredientPackageDto)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, packageOid
}
