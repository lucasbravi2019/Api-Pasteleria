package factory

import (
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/handlers"
	"github.com/lucasbravi2019/pasteleria/services"
)

func GetRecipeIngredientHandlerInstance() *handlers.RecipeIngredientHandler {
	if handlers.RecipeIngredientHandlerInstance == nil {
		handlers.RecipeIngredientHandlerInstance = &handlers.RecipeIngredientHandler{
			Service: *GetRecipeIngredientServiceInstance(),
		}
	}
	return handlers.RecipeIngredientHandlerInstance
}

func GetRecipeIngredientServiceInstance() *services.RecipeIngredientService {
	if services.RecipeIngredientServiceInstance == nil {
		services.RecipeIngredientServiceInstance = &services.RecipeIngredientService{
			RecipeDao:           *GetRecipeDaoInstance(),
			IngredientDao:       *GetIngredientDaoInstance(),
			RecipeIngredientDao: *GetRecipeIngredientDaoInstance(),
		}
	}
	return services.RecipeIngredientServiceInstance
}

func GetRecipeIngredientDaoInstance() *dao.RecipeIngredientDao {
	if dao.RecipeIngredientDaoInstance == nil {
		dao.RecipeIngredientDaoInstance = &dao.RecipeIngredientDao{
			DB: core.GetDatabaseConnection(),
		}
	}
	return dao.RecipeIngredientDaoInstance
}
