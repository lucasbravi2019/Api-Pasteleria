package services

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/pkg"
	"github.com/lucasbravi2019/pasteleria/pkg/util"
)

type IngredientService struct {
	IngredientDao *dao.IngredientDao
	RecipeService *RecipeService
}

var IngredientServiceInstance *IngredientService

func (s *IngredientService) GetAllIngredients() (int, *[]dto.IngredientResponse, error) {
	ingredients, err := s.IngredientDao.GetAllIngredients()

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusOK, ingredients, nil
}

func (s *IngredientService) CreateIngredient(ctx *gin.Context) (int, interface{}, error) {
	var ingredient dto.IngredientRequest

	err := pkg.DecodeBody(ctx, &ingredient)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.CreateIngredientName(ingredient.Name)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	ingredientId, err := s.IngredientDao.FindIngredientIdByName(ingredient.Name)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	err = s.IngredientDao.AddIngredientPackage(ingredientId, ingredient.Packages)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	return http.StatusCreated, nil, nil
}

func (s *IngredientService) UpdateIngredient(ctx *gin.Context) (int, interface{}, error) {
	var ingredient dto.IngredientRequest

	err := pkg.DecodeBody(ctx, &ingredient)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.UpdateIngredientName(&ingredient)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	ids := s.getPackagesIdsToUpdateOrInsert(&ingredient)
	idsExistent, err := s.IngredientDao.FindPackagesIdByIngredientId(ingredient.Id)

	if pkg.HasError(err) {
		log.Println("despues de getPackagesIdsToUpdateOrInsert")
		return http.StatusInternalServerError, nil, err
	}

	idsToRemove := s.getPackagesIdsToRemove(ids, idsExistent)

	if len(*idsToRemove) > 0 {
		err = s.IngredientDao.RemoveIngredientPackages(idsToRemove)

		if pkg.HasError(err) {
			log.Println("despues de getPackagesIdsToRemove")
			return http.StatusInternalServerError, nil, err
		}
	}

	if ingredient.Packages != nil && len(*ingredient.Packages) > 0 {
		err = s.IngredientDao.AddIngredientPackage(ingredient.Id, ingredient.Packages)

		if pkg.HasError(err) {
			log.Println("despues de AddIngredientPackage")
			return http.StatusInternalServerError, nil, err
		}
	}

	_, recipes, err := s.RecipeService.GetAllRecipes()

	if pkg.HasError(err) {
		log.Println("despues de getallrecipes")
		return http.StatusInternalServerError, nil, err
	}

	for _, recipe := range *recipes {
		err := s.RecipeService.UpdateRecipePrice(&recipe)

		if pkg.HasError(err) {
			return http.StatusInternalServerError, nil, err
		}
	}

	return http.StatusOK, nil, nil
}

func (s *IngredientService) DeleteIngredient(ctx *gin.Context) (int, interface{}, error) {
	id, err := pkg.GetUrlVars(ctx, "id")

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	ingredientId, err := util.ToLong(id)

	if pkg.HasError(err) {
		return http.StatusBadRequest, nil, err
	}

	err = s.IngredientDao.DeleteIngredient(&ingredientId)

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	_, recipes, err := s.RecipeService.GetAllRecipes()

	if pkg.HasError(err) {
		return http.StatusInternalServerError, nil, err
	}

	for _, recipe := range *recipes {
		err := s.RecipeService.UpdateRecipePrice(&recipe)

		if pkg.HasError(err) {
			return http.StatusInternalServerError, nil, err
		}
	}

	return http.StatusOK, nil, nil
}

func (s *IngredientService) getPackagesIdsToUpdateOrInsert(request *dto.IngredientRequest) *[]int64 {
	packagesIds := util.NewList[int64]()
	if request.Packages == nil {
		return &packagesIds
	}

	for _, pkg := range *request.Packages {
		if pkg.Id != nil {
			util.Add(&packagesIds, *pkg.Id)
		}
	}

	return &packagesIds
}

func (s *IngredientService) getPackagesIdsToRemove(requestIds *[]int64, existentIds *[]int64) *[]int64 {
	idsToRemove := util.NewList[int64]()
	if requestIds == nil || existentIds == nil {
		return &idsToRemove
	}

	for _, existentId := range *existentIds {
		found := false
		for _, requestId := range *requestIds {
			if existentId == requestId {
				found = true
			}
		}

		if !found {
			util.Add(&idsToRemove, existentId)
		}
	}

	return &idsToRemove
}
