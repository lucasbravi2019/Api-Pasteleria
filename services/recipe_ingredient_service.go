package services

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeIngredientService struct {
	RecipeDao           dao.RecipeDao
	IngredientDao       dao.IngredientDao
	RecipeIngredientDao dao.RecipeIngredientDao
}

type RecipeIngredientServiceInterface interface {
	AddIngredientToRecipe(r *http.Request) int
}

var RecipeIngredientServiceInstance *RecipeIngredientService

func (s *RecipeIngredientService) AddIngredientToRecipe(r *http.Request) int {
	recipeOid := core.ConvertHexToObjectId(mux.Vars(r)["recipeId"])
	ingredientOid := core.ConvertHexToObjectId(mux.Vars(r)["ingredientId"])

	if recipeOid == nil || ingredientOid == nil {
		return http.StatusBadRequest
	}

	recipe := s.RecipeDao.FindRecipeByOID(recipeOid)

	if recipe == nil {
		return http.StatusNotFound
	}

	ingredientDTO := s.IngredientDao.FindIngredientByOID(ingredientOid)

	if ingredientDTO == nil {
		return http.StatusNotFound
	}

	ingredientDetails := &dto.IngredientDetailsDTO{}

	invalidBody := core.DecodeBody(r, ingredientDetails)

	if invalidBody {
		return http.StatusBadRequest
	}

	err := validate(ingredientDTO, ingredientDetails)

	if err != nil {
		return http.StatusBadRequest
	}

	envase := getIngredientPackage(ingredientDetails.Metric, ingredientDTO.Packages)

	recipeIngredient := &models.RecipeIngredient{
		ID:       primitive.NewObjectID(),
		Quantity: ingredientDetails.Quantity,
		Name:     ingredientDTO.Name,
		Package: models.RecipeIngredientPackage{
			ID:       envase.ID,
			Metric:   envase.Metric,
			Quantity: envase.Quantity,
			Price:    envase.Price,
		},
		Price: float64(ingredientDetails.Quantity) / envase.Quantity * envase.Price,
	}

	err = s.RecipeIngredientDao.AddIngredientToRecipe(recipeOid, recipeIngredient)

	if err != nil {
		return http.StatusInternalServerError
	}

	err = s.RecipeDao.UpdateRecipeByIdPrice(recipeOid)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError
	}

	return http.StatusOK
}
