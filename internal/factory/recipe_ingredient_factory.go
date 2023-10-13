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
			RecipeDao:           *GetRecipeDaoInstance(),
			IngredientDao:       *GetIngredientDaoInstance(),
			RecipeIngredientDao: *GetRecipeIngredientDaoInstance(),
			RecipeMapper:        *GetRecipeMapperInstance(),
		}
	}
	return services.RecipeIngredientServiceInstance
}

func GetRecipeIngredientDaoInstance() *dao.RecipeIngredientDao {
	if dao.RecipeIngredientDaoInstance == nil {
		dao.RecipeIngredientDaoInstance = &dao.RecipeIngredientDao{
			DB: db.GetDatabaseConnection(),
		}
	}
	return dao.RecipeIngredientDaoInstance
}

func GetRecipeMapperInstance() *mapper.RecipeMapper {
	if mapper.RecipeMapperInstance == nil {
		mapper.RecipeMapperInstance = &mapper.RecipeMapper{}
	}
	return mapper.RecipeMapperInstance
}
