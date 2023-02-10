package ingredients

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/packages"
	"github.com/lucasbravi2019/pasteleria/core"
)

type service struct {
	repository IngredientRepository
}

type IngredientService interface {
	GetAllIngredients() (int, []IngredientDTO)
	CreateIngredient(r *http.Request) (int, *IngredientDTO)
	UpdateIngredient(r *http.Request) (int, *IngredientDTO)
	DeleteIngredient(r *http.Request) (int, *IngredientDTO)
	AddPackageToIngredient(r *http.Request) (int, *IngredientDTO)
	ChangeIngredientPrice(r *http.Request) (int, *IngredientDTO)
}

var ingredientServiceInstance *service

func (s *service) GetAllIngredients() (int, []IngredientDTO) {
	return s.repository.GetAllIngredients()
}

func (s *service) CreateIngredient(r *http.Request) (int, *IngredientDTO) {
	var ingredient *IngredientDTO = &IngredientDTO{}

	invalidBody := core.DecodeBody(r, ingredient)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	ingredient.Package = []packages.Package{}

	return s.repository.CreateIngredient(ingredient)
}

func (s *service) UpdateIngredient(r *http.Request) (int, *IngredientDTO) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	var ingredient *IngredientDTO = &IngredientDTO{}

	invalidBody := core.DecodeBody(r, ingredient)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.repository.UpdateIngredient(oid, ingredient)
}

func (s *service) DeleteIngredient(r *http.Request) (int, *IngredientDTO) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.repository.DeleteIngredient(oid)
}

func (s *service) AddPackageToIngredient(r *http.Request) (int, *IngredientDTO) {
	ingredientOid := mux.Vars(r)["ingredientId"]
	packageOid := mux.Vars(r)["packageId"]

	var priceDTO *IngredientPackagePriceDTO = &IngredientPackagePriceDTO{}

	core.DecodeBody(r, priceDTO)

	var ingredientPackageDto *IngredientPackageDTO = &IngredientPackageDTO{
		IngredientOid: *core.ConvertHexToObjectId(ingredientOid),
		PackageOid:    *core.ConvertHexToObjectId(packageOid),
		Price:         priceDTO.Price,
	}

	return s.repository.AddPackageToIngredient(*ingredientPackageDto)
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

	return s.repository.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)
}
