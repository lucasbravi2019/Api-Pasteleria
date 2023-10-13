package services

import (
	"net/http"

	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
)

type RecipeService struct {
	RecipeDao dao.RecipeDao
}

type RecipeServiceInterface interface {
	GetAllRecipes() (int, *[]dto.RecipeDTO)
	GetRecipe(r *http.Request) (int, *dto.RecipeDTO)
	CreateRecipe(r *http.Request) error
	UpdateRecipeName(r *http.Request) int
	DeleteRecipe(r *http.Request) int
}

var RecipeServiceInstance *RecipeService

func (s *RecipeService) GetAllRecipes() (int, *[]dto.RecipeDTO, error) {

	return http.StatusOK, nil, nil
}

func (s *RecipeService) GetRecipe(r *http.Request) (int, *dto.RecipeDTO) {
	return 200, nil
}

func (s *RecipeService) CreateRecipe(r *http.Request) error {

	return nil
}

func (s *RecipeService) UpdateRecipeName(r *http.Request) int {

	return http.StatusOK
}

func (s *RecipeService) DeleteRecipe(r *http.Request) int {

	return http.StatusOK
}
