package factory

import (
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/handlers"
	"github.com/lucasbravi2019/pasteleria/services"
)

func GetIngredientHandlerInstance() *handlers.IngredientHandler {
	if handlers.IngredientHandlerInstance == nil {
		handlers.IngredientHandlerInstance = &handlers.IngredientHandler{
			Service: GetIngredientServiceInstance(),
		}
	}
	return handlers.IngredientHandlerInstance
}

func GetIngredientServiceInstance() *services.IngredientService {
	if services.IngredientServiceInstance == nil {
		services.IngredientServiceInstance = &services.IngredientService{
			IngredientDao:       *GetIngredientDaoInstance(),
			RecipeDao:           *GetRecipeDaoInstance(),
			RecipeIngredientDao: *GetRecipeIngredientDaoInstance(),
		}
	}
	return services.IngredientServiceInstance
}

func GetIngredientDaoInstance() *dao.IngredientDao {
	if dao.IngredientDaoInstance == nil {
		dao.IngredientDaoInstance = &dao.IngredientDao{
			DB: core.GetDatabaseConnection(),
		}
	}
	return dao.IngredientDaoInstance
}
