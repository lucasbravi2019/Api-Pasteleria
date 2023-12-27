package factory

import (
	"github.com/lucasbravi2019/pasteleria/api"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/services"
)

func GetRecipeIngredientHandlerInstance() *api.RecipeIngredientHandler {
	if api.RecipeIngredientHandlerInstance == nil {
		api.RecipeIngredientHandlerInstance = &api.RecipeIngredientHandler{
			Service: *GetRecipeIngredientServiceInstance(),
		}
	}
	return api.RecipeIngredientHandlerInstance
}

func GetRecipeIngredientServiceInstance() *services.RecipeIngredientService {
	if services.RecipeIngredientServiceInstance == nil {
		services.RecipeIngredientServiceInstance = &services.RecipeIngredientService{
			RecipeIngredientDao: *GetRecipeIngredientDaoInstance(),
		}
	}
	return services.RecipeIngredientServiceInstance
}

func GetRecipeIngredientDaoInstance() *dao.RecipeIngredientDao {
	if dao.RecipeIngredientDaoInstance == nil {
		dao.RecipeIngredientDaoInstance = &dao.RecipeIngredientDao{
			DB:                     db.GetDatabaseConnection(),
			RecipeIngredientMapper: GetRecipeIngredientMapperInstance(),
		}
	}
	return dao.RecipeIngredientDaoInstance
}

func GetRecipeIngredientMapperInstance() *mapper.RecipeIngredientMapper {
	if mapper.RecipeIngredientMapperInstance == nil {
		mapper.RecipeIngredientMapperInstance = &mapper.RecipeIngredientMapper{}
	}
	return mapper.RecipeIngredientMapperInstance
}
