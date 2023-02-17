package services

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/mapper"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeIngredientService struct {
	RecipeDao           dao.RecipeDao
	IngredientDao       dao.IngredientDao
	RecipeIngredientDao dao.RecipeIngredientDao
	RecipeMapper        mapper.RecipeMapper
}

type RecipeIngredientServiceInterface interface {
	AddIngredientToRecipe(r *http.Request) int
	RemoveIngredientFromRecipe(r *http.Request) (int, *dto.RecipeDTO)
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

func (s *RecipeIngredientService) RemoveIngredientFromRecipe(r *http.Request) (int, *dto.RecipeDTO) {
	ids := dto.RecipeIngredientIdDTO{}

	invalidBody := core.DecodeBody(r, &ids)
	log.Println(ids)

	if invalidBody {
		return http.StatusBadRequest, nil
	}

	recipeId := core.ConvertHexToObjectId(ids.RecipeId)
	ingredientId := core.ConvertHexToObjectId(ids.IngredientId)

	if recipeId == nil || ingredientId == nil {
		return http.StatusInternalServerError, nil
	}

	originalRecipe := s.RecipeDao.FindRecipeByOID(recipeId)

	if originalRecipe == nil {
		return http.StatusNotFound, nil
	}

	for i := 0; i < len(originalRecipe.Ingredients); i++ {
		if originalRecipe.Ingredients[i].ID == *ingredientId {
			originalRecipe.Ingredients = append(originalRecipe.Ingredients[:i], originalRecipe.Ingredients[i+1:]...)
		}
	}

	recipeEntity := s.RecipeMapper.RecipeDTOToRecipe(originalRecipe)

	err := s.RecipeIngredientDao.RemoveIngredientFromRecipe(recipeId, recipeEntity)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	err = s.RecipeDao.UpdateRecipesPrice()

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, originalRecipe
}
