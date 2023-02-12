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
	GetPackages() (int, []Package)
	CreatePackage(r *http.Request) (int, *Package)
	UpdatePackage(r *http.Request) (int, *Package)
	DeletePackage(r *http.Request) (int, *Package)
	AddPackageToIngredient(r *http.Request) (int, *ingredients.IngredientDTO)
	RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID)
}

var packageServiceInstance *service

func (s *service) GetPackages() (int, []Package) {
	return s.packageRepository.GetPackages()
}

func (s *service) CreatePackage(r *http.Request) (int, *Package) {
	var packageRequest *Package = &Package{}

	invalidBody := core.DecodeBody(r, packageRequest)

	if invalidBody {
		return http.StatusBadRequest, nil
	}
	return s.packageRepository.CreatePackage(packageRequest)
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

	return s.packageRepository.UpdatePackage(oid, packageRequest)
}

func (s *service) DeletePackage(r *http.Request) (int, *Package) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}
	_, envase := s.packageRepository.DeletePackage(oid)

	var ingredientPackage *ingredients.IngredientPackageDTO = &ingredients.IngredientPackageDTO{
		PackageOid: *oid,
	}
	s.ingredientRepository.RemovePackageFromIngredients(*ingredientPackage)
	s.recipeRepository.RemoveIngredientByPackageId(oid)
	s.recipeRepository.UpdateRecipesPrice()

	return http.StatusOK, envase
}

func (s *service) AddPackageToIngredient(r *http.Request) (int, *ingredients.IngredientDTO) {
	ingredientOid := mux.Vars(r)["ingredientId"]
	packageOid := mux.Vars(r)["packageId"]
	ingredientId := core.ConvertHexToObjectId(ingredientOid)
	packageId := core.ConvertHexToObjectId(packageOid)

	var priceDTO *ingredients.IngredientPackagePriceDTO = &ingredients.IngredientPackagePriceDTO{}

	invalidBody := core.DecodeBody(r, priceDTO)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	_, envase := s.packageRepository.GetPackageById(packageId)

	var ingredientPackage *ingredients.IngredientPackage = &ingredients.IngredientPackage{
		ID:       envase.ID,
		Metric:   envase.Metric,
		Quantity: envase.Quantity,
		Price:    priceDTO.Price,
	}

	return s.ingredientRepository.AddPackageToIngredient(ingredientId, packageId, ingredientPackage)
}

func (s *service) RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID) {
	packageOid := mux.Vars(r)["packageId"]

	var ingredientPackageDto *ingredients.IngredientPackageDTO = &ingredients.IngredientPackageDTO{
		PackageOid: *core.ConvertHexToObjectId(packageOid),
	}

	return s.ingredientRepository.RemovePackageFromIngredients(*ingredientPackageDto)
}
