package packages

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/ingredients"
	"github.com/lucasbravi2019/pasteleria/api/recipes"
	"github.com/lucasbravi2019/pasteleria/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	packageRepository    PackageRepository
	ingredientRepository ingredients.IngredientRepository
	recipeRepository     recipes.RecipeRepository
}

type PackageService interface {
	GetPackages() (int, *[]Package)
	CreatePackage(r *http.Request) (int, *Package)
	UpdatePackage(r *http.Request) (int, *Package)
	DeletePackage(r *http.Request) (int, *primitive.ObjectID)
	AddPackageToIngredient(r *http.Request) int
	RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID)
}

var packageServiceInstance *service

func (s *service) GetPackages() (int, *[]Package) {
	packages := s.packageRepository.GetPackages()

	if packages == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, packages
}

func (s *service) CreatePackage(r *http.Request) (int, *Package) {
	var packageRequest *Package = &Package{}

	invalidBody := core.DecodeBody(r, packageRequest)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	id := s.packageRepository.CreatePackage(packageRequest)

	if id == nil {
		return http.StatusInternalServerError, nil
	}

	envase := s.packageRepository.GetPackageById(id)

	if envase == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusCreated, envase
}

func (s *service) UpdatePackage(r *http.Request) (int, *Package) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	var packageRequest *Package = &Package{}

	invalidBody := core.DecodeBody(r, packageRequest)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	err := s.packageRepository.UpdatePackage(oid, packageRequest)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	envase := s.packageRepository.GetPackageById(oid)

	if envase == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, envase
}

func (s *service) DeletePackage(r *http.Request) (int, *primitive.ObjectID) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	err := s.packageRepository.DeletePackage(oid)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	var ingredientPackage *ingredients.IngredientPackageDTO = &ingredients.IngredientPackageDTO{
		PackageOid: *oid,
	}

	err = s.ingredientRepository.RemovePackageFromIngredients(*ingredientPackage)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	err = s.recipeRepository.RemoveIngredientByPackageId(oid)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	err = s.recipeRepository.UpdateRecipesPrice()

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, oid
}

func (s *service) AddPackageToIngredient(r *http.Request) int {
	ingredientOid := mux.Vars(r)["ingredientId"]
	packageOid := mux.Vars(r)["packageId"]
	ingredientId := core.ConvertHexToObjectId(ingredientOid)
	packageId := core.ConvertHexToObjectId(packageOid)

	var priceDTO *ingredients.IngredientPackagePriceDTO = &ingredients.IngredientPackagePriceDTO{}

	invalidBody := core.DecodeBody(r, priceDTO)

	if invalidBody {
		return http.StatusBadRequest
	}

	envase := s.packageRepository.GetPackageById(packageId)

	if envase == nil {
		return http.StatusNotFound
	}

	var ingredientPackage *ingredients.IngredientPackage = &ingredients.IngredientPackage{
		ID:       envase.ID,
		Metric:   envase.Metric,
		Quantity: envase.Quantity,
		Price:    priceDTO.Price,
	}

	err := s.ingredientRepository.AddPackageToIngredient(ingredientId, packageId, ingredientPackage)

	if err != nil {
		return http.StatusInternalServerError
	}

	return http.StatusOK
}

func (s *service) RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID) {
	packageOid := core.ConvertHexToObjectId(mux.Vars(r)["packageId"])

	var ingredientPackageDto *ingredients.IngredientPackageDTO = &ingredients.IngredientPackageDTO{
		PackageOid: *packageOid,
	}

	err := s.ingredientRepository.RemovePackageFromIngredients(*ingredientPackageDto)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, packageOid
}
