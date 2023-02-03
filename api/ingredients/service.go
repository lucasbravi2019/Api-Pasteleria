package ingredients

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
)

type ingredientService struct {
	repository IngredientRepository
}

type IngredientService interface {
	GetAllIngredients() (int, []Ingredient)
	CreateIngredient(r *http.Request) (int, *Ingredient)
	UpdateIngredient(r *http.Request) (int, *Ingredient)
	DeleteIngredient(r *http.Request) (int, *Ingredient)
}

var ingredientServiceInstance *ingredientService

func (s *ingredientService) GetAllIngredients() (int, []Ingredient) {
	return s.repository.GetAllIngredients()
}

func (s *ingredientService) CreateIngredient(r *http.Request) (int, *Ingredient) {
	var ingredient *Ingredient = &Ingredient{}

	invalidBody := core.DecodeBody(r, ingredient)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	return s.repository.CreateIngredient(ingredient)
}

func (s *ingredientService) UpdateIngredient(r *http.Request) (int, *Ingredient) {
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

func (s *ingredientService) DeleteIngredient(r *http.Request) (int, *Ingredient) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	return s.repository.DeleteIngredient(oid)
}
