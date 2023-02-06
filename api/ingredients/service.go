package ingredients

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
)

type service struct {
	repository IngredientRepository
}

type IngredientService interface {
	GetAllIngredients() (int, []Ingredient)
	CreateIngredient(r *http.Request) (int, *Ingredient)
	UpdateIngredient(r *http.Request) (int, *Ingredient)
	DeleteIngredient(r *http.Request) (int, *Ingredient)
	AddPackageToIngredient(r *http.Request) (int, *Ingredient)
	ChangeIngredientPrice(r *http.Request) (int, *[]Ingredient)
}

var ingredientServiceInstance *service

func (s *service) GetAllIngredients() (int, []Ingredient) {
	return s.repository.GetAllIngredients()
}

func (s *service) CreateIngredient(r *http.Request) (int, *Ingredient) {
	var ingredient *Ingredient = &Ingredient{}

	invalidBody := core.DecodeBody(r, ingredient)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.repository.CreateIngredient(ingredient)
}

func (s *service) UpdateIngredient(r *http.Request) (int, *Ingredient) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	var ingredient *Ingredient = &Ingredient{}

	invalidBody := core.DecodeBody(r, ingredient)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.repository.UpdateIngredient(oid, ingredient)
}

func (s *service) DeleteIngredient(r *http.Request) (int, *Ingredient) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.repository.DeleteIngredient(oid)
}

func (s *service) AddPackageToIngredient(r *http.Request) (int, *Ingredient) {
	ingredientOid := mux.Vars(r)["ingredientId"]
	packageOid := mux.Vars(r)["packageId"]

	var ingredientPackagePrice *IngredientPackagePrice = &IngredientPackagePrice{}

	core.DecodeBody(r, ingredientPackagePrice)

	var ingredientPackageDto *IngredientPackageDTO = &IngredientPackageDTO{
		IngredientOid: *core.ConvertHexToObjectId(ingredientOid),
		PackageOid:    *core.ConvertHexToObjectId(packageOid),
		Price:         ingredientPackagePrice.Price,
	}

	return s.repository.AddPackageToIngredient(*ingredientPackageDto)
}

func (s *service) ChangeIngredientPrice(r *http.Request) (int, *[]Ingredient) {
	ingredientPackageId := mux.Vars(r)["id"]
	ingredientPackageOid := core.ConvertHexToObjectId(ingredientPackageId)

	if ingredientPackageOid == nil {
		return http.StatusBadRequest, nil
	}

	var ingredientPackagePrice *IngredientPackagePrice = &IngredientPackagePrice{}

	invalidBody := core.DecodeBody(r, ingredientPackagePrice)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.repository.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)
}
