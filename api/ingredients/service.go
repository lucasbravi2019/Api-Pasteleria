package ingredients

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/packages"
	"github.com/lucasbravi2019/pasteleria/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	ingredientRepository IngredientRepository
	packageRepository    packages.PackageRepository
}

type IngredientService interface {
	GetAllIngredients() (int, []IngredientDTO)
	CreateIngredient(r *http.Request) (int, *IngredientDTO)
	UpdateIngredient(r *http.Request) (int, *IngredientDTO)
	DeleteIngredient(r *http.Request) (int, *primitive.ObjectID)
	AddPackageToIngredient(r *http.Request) (int, *IngredientDTO)
	RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID)
	ChangeIngredientPrice(r *http.Request) (int, *IngredientDTO)
}

var ingredientServiceInstance *service

func (s *service) GetAllIngredients() (int, []IngredientDTO) {
	return s.ingredientRepository.GetAllIngredients()
}

func (s *service) CreateIngredient(r *http.Request) (int, *IngredientDTO) {
	var ingredientDto *IngredientNameDTO = &IngredientNameDTO{}

	invalidBody := core.DecodeBody(r, ingredientDto)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.ingredientRepository.CreateIngredient(ingredientDto)
}

func (s *service) UpdateIngredient(r *http.Request) (int, *IngredientDTO) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	var ingredient *IngredientNameDTO = &IngredientNameDTO{}

	invalidBody := core.DecodeBody(r, ingredient)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.ingredientRepository.UpdateIngredient(oid, ingredient)
}

func (s *service) DeleteIngredient(r *http.Request) (int, *primitive.ObjectID) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.ingredientRepository.DeleteIngredient(oid)
}

func (s *service) AddPackageToIngredient(r *http.Request) (int, *IngredientDTO) {
	ingredientOid := mux.Vars(r)["ingredientId"]
	packageOid := mux.Vars(r)["packageId"]
	ingredientId := core.ConvertHexToObjectId(ingredientOid)
	packageId := core.ConvertHexToObjectId(packageOid)

	var priceDTO *IngredientPackagePriceDTO = &IngredientPackagePriceDTO{}

	invalidBody := core.DecodeBody(r, priceDTO)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	_, envase := s.packageRepository.GetPackageById(packageId)

	var ingredientPackage *IngredientPackage = &IngredientPackage{
		ID:       envase.ID,
		Metric:   envase.Metric,
		Quantity: envase.Quantity,
		Price:    priceDTO.Price,
	}

	return s.ingredientRepository.AddPackageToIngredient(ingredientId, packageId, ingredientPackage)
}

func (s *service) RemovePackageFromIngredients(r *http.Request) (int, *primitive.ObjectID) {
	packageOid := mux.Vars(r)["packageId"]

	var ingredientPackageDto *IngredientPackageDTO = &IngredientPackageDTO{
		PackageOid: *core.ConvertHexToObjectId(packageOid),
	}

	return s.ingredientRepository.RemovePackageFromIngredients(*ingredientPackageDto)
}

func (s *service) ChangeIngredientPrice(r *http.Request) (int, *IngredientDTO) {
	ingredientPackageId := mux.Vars(r)["id"]
	ingredientPackageOid := core.ConvertHexToObjectId(ingredientPackageId)

	if ingredientPackageOid == nil {
		return http.StatusBadRequest, nil
	}

	var ingredientPackagePrice *IngredientPackagePriceDTO = &IngredientPackagePriceDTO{}

	invalidBody := core.DecodeBody(r, ingredientPackagePrice)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.ingredientRepository.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)
}
