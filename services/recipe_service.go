package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeService struct {
	RecipeDao dao.RecipeDao
}

type RecipeServiceInterface interface {
	GetAllRecipes() (int, *[]dto.RecipeDTO, error)
	GetRecipe(c *gin.Context) (int, *dto.RecipeDTO, error)
	CreateRecipe(c *gin.Context) (int, *dto.RecipeDTO, error)
	UpdateRecipeName(c *gin.Context) (int, *dto.RecipeDTO, error)
	DeleteRecipe(c *gin.Context) (int, *primitive.ObjectID, error)
}

var RecipeServiceInstance *RecipeService

func (s *RecipeService) GetAllRecipes() (int, *[]dto.RecipeDTO, error) {
	recipes, err := s.RecipeDao.FindAllRecipes()

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, recipes, nil
}

func (s *RecipeService) GetRecipe(c *gin.Context) (int, *dto.RecipeDTO, error) {
	recipeId, err := core.ConvertUrlParamToObjectId("id", c)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	recipe, err := s.RecipeDao.FindRecipeByOID(recipeId)

	if err != nil {
		return http.StatusNotFound, nil, err
	}

	return http.StatusOK, recipe, nil
}

func (s *RecipeService) CreateRecipe(c *gin.Context) (int, *dto.RecipeDTO, error) {
	recipeName := &dto.RecipeNameDTO{}

	err := core.DecodeBody(c, recipeName)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	recipeEntity := &models.Recipe{
		Name:        recipeName.Name,
		Ingredients: []models.RecipeIngredient{},
	}

	oid, err := s.RecipeDao.CreateRecipe(recipeEntity)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	recipeCreated, err := s.RecipeDao.FindRecipeByOID(oid)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, recipeCreated, nil
}

func (s *RecipeService) UpdateRecipeName(c *gin.Context) (int, *dto.RecipeDTO, error) {
	recipeName := &dto.RecipeNameDTO{}

	err := core.DecodeBody(c, recipeName)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	oid, err := core.ConvertToObjectId(recipeName.ID)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	recipe := &models.Recipe{
		Name: recipeName.Name,
	}

	err = s.RecipeDao.UpdateRecipeName(oid, recipe)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	recipeUpdated, err := s.RecipeDao.FindRecipeByOID(oid)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, recipeUpdated, nil
}

func (s *RecipeService) DeleteRecipe(c *gin.Context) (int, *primitive.ObjectID, error) {
	oid, err := core.ConvertUrlVarToObjectId("id", c)

	if err != nil {
		return http.StatusBadRequest, nil, err
	}

	log.Println(oid)

	err = s.RecipeDao.DeleteRecipe(oid)

	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, oid, nil
}
