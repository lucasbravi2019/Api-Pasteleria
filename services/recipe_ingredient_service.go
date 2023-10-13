package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/mapper"
	"github.com/lucasbravi2019/pasteleria/models"
)

type RecipeIngredientService struct {
	RecipeDao           dao.RecipeDao
	IngredientDao       dao.IngredientDao
	RecipeIngredientDao dao.RecipeIngredientDao
	RecipeMapper        mapper.RecipeMapper
}

type RecipeIngredientServiceInterface interface {
	AddIngredientToRecipe(c *gin.Context) (int, error)
	RemoveIngredientFromRecipe(c *gin.Context) (int, *dto.RecipeDTO, error)
}

var RecipeIngredientServiceInstance *RecipeIngredientService

func (s *RecipeIngredientService) AddIngredientToRecipe(c *gin.Context) (int, error) {
	recipeOid, err := core.ConvertUrlVarToObjectId("recipeId", c)

	if err != nil {
		return http.StatusBadRequest, err
	}

	ingredientOid, err := core.ConvertUrlVarToObjectId("ingredientId", c)

	if err != nil {
		return http.StatusBadRequest, err
	}

	_, err = s.RecipeDao.FindRecipeByOID(recipeOid)

	if err != nil {
		return http.StatusNotFound, err
	}

	ingredientDTO, err := s.IngredientDao.FindIngredientByOID(ingredientOid)

	if err != nil {
		return http.StatusNotFound, err
	}

	ingredientDetails := &dto.IngredientDetailsDTO{}

	err = core.DecodeBody(c, ingredientDetails)

	if err != nil {
		return http.StatusBadRequest, err
	}

	err = validate(ingredientDTO, ingredientDetails)

	if err != nil {
		return http.StatusBadRequest, err
	}

	envase := getIngredientPackage(ingredientDetails.Metric, ingredientDTO.Packages)

	recipeIngredient := &models.RecipeIngredient{
		// ID:       primitive.NewObjectID(),
		// Quantity: ingredientDetails.Quantity,
		Name: ingredientDTO.Name,
		Package: models.RecipeIngredientPackage{
			// ID:       envase.ID,
			Metric:   envase.Metric,
			Quantity: envase.Quantity,
			Price:    envase.Price,
		},
		Price: float64(ingredientDetails.Quantity) / envase.Quantity * envase.Price,
	}

	err = s.RecipeIngredientDao.AddIngredientToRecipe(recipeOid, recipeIngredient)

	if err != nil {
		return http.StatusInternalServerError, err
	}

	err = s.RecipeDao.UpdateRecipeByIdPrice(recipeOid)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, err
	}

	return http.StatusOK, nil
}

func (s *RecipeIngredientService) RemoveIngredientFromRecipe(c *gin.Context) (int, *dto.RecipeDTO, error) {
	ids := dto.RecipeIngredientIdDTO{}

	err := core.DecodeBody(c, &ids)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	recipeId, err := core.ConvertToObjectId(ids.RecipeId)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	ingredientId, err := core.ConvertToObjectId(ids.IngredientId)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	originalRecipe, err := s.RecipeDao.FindRecipeByOID(recipeId)

	if err != nil {
		return http.StatusNotFound, nil, err
	}

	for i := 0; i < len(originalRecipe.Ingredients); i++ {
		if originalRecipe.Ingredients[i].ID == *ingredientId {
			originalRecipe.Ingredients = append(originalRecipe.Ingredients[:i], originalRecipe.Ingredients[i+1:]...)
		}
	}

	recipeEntity := s.RecipeMapper.RecipeDTOToRecipe(originalRecipe)

	err = s.RecipeIngredientDao.RemoveIngredientFromRecipe(recipeId, recipeEntity)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	err = s.RecipeDao.UpdateRecipesPrice()

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, originalRecipe, nil
}
