package factory

import (
	"log"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/handlers"
	"github.com/lucasbravi2019/pasteleria/services"
)

func GetRecipeHandlerInstance() *handlers.RecipeHandler {
	log.Println("Recipe handler")
	if handlers.RecipeHandlerInstance == nil {
		handlers.RecipeHandlerInstance = &handlers.RecipeHandler{
			Service: *GetRecipeServiceInstance(),
		}
	}
	return handlers.RecipeHandlerInstance
}

func GetRecipeServiceInstance() *services.RecipeService {
	if services.RecipeServiceInstance == nil {
		services.RecipeServiceInstance = &services.RecipeService{
			RecipeDao: *GetRecipeDaoInstance(),
		}
	}
	return services.RecipeServiceInstance
}

func GetRecipeDaoInstance() *dao.RecipeDao {
	if dao.RecipeDaoInstance == nil {
		dao.RecipeDaoInstance = &dao.RecipeDao{
			DB: core.GetDatabaseConnection(),
		}
	}
	return dao.RecipeDaoInstance
}
