package ingredients

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/api/recipes"
	"github.com/lucasbravi2019/pasteleria/core"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type service struct {
	ingredientRepository IngredientRepository
	recipeRepository     recipes.RecipeRepository
}

type IngredientService interface {
	GetAllIngredients() (int, []IngredientDTO)
	CreateIngredient(r *http.Request) (int, *IngredientDTO)
	UpdateIngredient(r *http.Request) (int, *IngredientDTO)
	DeleteIngredient(r *http.Request) (int, *primitive.ObjectID)
	AddIngredientToRecipe(r *http.Request) (int, *primitive.ObjectID)
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

func (s *service) AddIngredientToRecipe(r *http.Request) (int, *primitive.ObjectID) {
	recipeOid := core.ConvertHexToObjectId(mux.Vars(r)["recipeId"])
	ingredientOid := core.ConvertHexToObjectId(mux.Vars(r)["ingredientId"])

	if recipeOid == nil || ingredientOid == nil {
		return http.StatusBadRequest, nil
	}

	_, recipe := s.recipeRepository.FindRecipeByOID(recipeOid)

	if recipe == nil {
		return http.StatusNotFound, nil
	}

	_, ingredientDTO := s.ingredientRepository.FindIngredientByOID(ingredientOid)

	if ingredientDTO == nil {
		return http.StatusNotFound, nil
	}

	var ingredientDetails *IngredientDetailsDTO = &IngredientDetailsDTO{}

	invalidBody := core.DecodeBody(r, ingredientDetails)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	err := validate(ingredientDTO, ingredientDetails)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	envase := getIngredientPackage(ingredientDetails.Metric, ingredientDTO.Packages)

	var recipeIngredient *recipes.RecipeIngredient = &recipes.RecipeIngredient{
		ID:       primitive.NewObjectID(),
		Quantity: ingredientDetails.Quantity,
		Name:     ingredientDTO.Name,
		Package: recipes.RecipeIngredientPackage{
			ID:       envase.ID,
			Metric:   envase.Metric,
			Quantity: envase.Quantity,
			Price:    envase.Price,
		},
		Price: float64(ingredientDetails.Quantity) / envase.Quantity * envase.Price,
	}

	_, recipeUpdated := s.recipeRepository.AddIngredientToRecipe(recipeOid, recipeIngredient)

	if recipeUpdated == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, nil
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

	_, dto := s.ingredientRepository.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)
	s.recipeRepository.UpdateIngredientsPrice(ingredientPackageOid, ingredientPackagePrice.Price)

	return http.StatusOK, dto
}

func validate(ingredient *IngredientDTO, ingredientDetails *IngredientDetailsDTO) error {
	if !ingredientMetricMatches(ingredientDetails.Metric, ingredient.Packages) {
		log.Println("La unidad de medida no coincide")
		return errors.New("la unidad de medida no coincide")
	}

	if ingredientDetails.Quantity == 0 {
		log.Println("La cantidad del ingrediente no puede ser 0")
		return errors.New("la cantidad del ingrediente no puede ser 0")
	}
	return nil
}

func ingredientMetricMatches(metric string, packages []PackageDTO) bool {
	for _, pack := range packages {
		if fmt.Sprintf("%g %s", pack.Quantity, pack.Metric) == metric {
			return true
		}
	}
	return false
}

func getIngredientPackage(metric string, packages []PackageDTO) *PackageDTO {
	for _, pack := range packages {
		if fmt.Sprintf("%g %s", pack.Quantity, pack.Metric) == metric {
			return &pack
		}
	}
	return nil
}
