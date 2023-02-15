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
	AddIngredientToRecipe(r *http.Request) int
	ChangeIngredientPrice(r *http.Request) (int, *IngredientDTO)
}

var ingredientServiceInstance *service

func (s *service) GetAllIngredients() (int, []IngredientDTO) {
	ingredients := s.ingredientRepository.GetAllIngredients()

	return http.StatusOK, ingredients
}

func (s *service) CreateIngredient(r *http.Request) (int, *IngredientDTO) {
	var ingredientDto *IngredientNameDTO = &IngredientNameDTO{}

	invalidBody := core.DecodeBody(r, ingredientDto)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	var ingredientEntity *Ingredient = &Ingredient{
		Name:     ingredientDto.Name,
		Packages: []IngredientPackage{},
	}

	ingredientCreatedId := s.ingredientRepository.CreateIngredient(ingredientEntity)

	if ingredientCreatedId == nil {
		return http.StatusInternalServerError, nil
	}

	ingredientCreated := s.ingredientRepository.FindIngredientByOID(ingredientCreatedId)

	if ingredientCreated == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusCreated, ingredientCreated
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

	err := s.ingredientRepository.UpdateIngredient(oid, ingredient)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	ingredientUpdated := s.ingredientRepository.FindIngredientByOID(oid)

	if ingredientUpdated == nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, ingredientUpdated
}

func (s *service) DeleteIngredient(r *http.Request) (int, *primitive.ObjectID) {
	oid := core.ConvertHexToObjectId(mux.Vars(r)["id"])

	if oid == nil {
		return http.StatusBadRequest, nil
	}

	err := s.ingredientRepository.DeleteIngredient(oid)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, oid
}

func (s *service) AddIngredientToRecipe(r *http.Request) int {
	recipeOid := core.ConvertHexToObjectId(mux.Vars(r)["recipeId"])
	ingredientOid := core.ConvertHexToObjectId(mux.Vars(r)["ingredientId"])

	if recipeOid == nil || ingredientOid == nil {
		return http.StatusBadRequest
	}

	recipe := s.recipeRepository.FindRecipeByOID(recipeOid)

	if recipe == nil {
		return http.StatusNotFound
	}

	ingredientDTO := s.ingredientRepository.FindIngredientByOID(ingredientOid)

	if ingredientDTO == nil {
		return http.StatusNotFound
	}

	var ingredientDetails *IngredientDetailsDTO = &IngredientDetailsDTO{}

	invalidBody := core.DecodeBody(r, ingredientDetails)

	if invalidBody {
		return http.StatusBadRequest
	}

	err := validate(ingredientDTO, ingredientDetails)

	if err != nil {
		return http.StatusBadRequest
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

	err = s.recipeRepository.AddIngredientToRecipe(recipeOid, recipeIngredient)

	if err != nil {
		return http.StatusInternalServerError
	}

	err = s.recipeRepository.UpdateRecipeByIdPrice(recipeOid)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError
	}

	return http.StatusOK
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

	err := s.ingredientRepository.ChangeIngredientPrice(ingredientPackageOid, ingredientPackagePrice)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	s.recipeRepository.UpdateIngredientPackagePrice(ingredientPackageOid, ingredientPackagePrice.Price)

	recipes := s.recipeRepository.FindRecipesByPackageId(ingredientPackageOid)

	if len(recipes) == 0 {
		return http.StatusInternalServerError, nil
	}

	for i := 0; i < len(recipes); i++ {
		var recipePrice float64 = 0
		for j := 0; j < len(recipes[i].Ingredients); j++ {
			recipes[i].Ingredients[j].Price = recipes[i].Ingredients[j].Quantity / recipes[i].Ingredients[j].Package.Quantity * recipes[i].Ingredients[j].Package.Price
			recipePrice += recipes[i].Ingredients[j].Price
		}
		recipes[i].Price = recipePrice

		err := s.recipeRepository.UpdateIngredientsPrice(ingredientPackageOid, recipes[i])

		if err != nil {
			log.Println(err.Error())
		}
	}

	ingredientUpdated := s.ingredientRepository.FindIngredientByPackageId(ingredientPackageOid)

	if ingredientUpdated == nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, ingredientUpdated
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
