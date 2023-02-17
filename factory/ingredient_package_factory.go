package factory

import (
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/handlers"
	"github.com/lucasbravi2019/pasteleria/services"
)

func GetIngredientPackageHandlerInstance() *handlers.IngredientPackageHandler {
	if handlers.IngredientPackageHandlerInstance == nil {
		handlers.IngredientPackageHandlerInstance = &handlers.IngredientPackageHandler{
			Service: *GetIngredientPackageServiceInstance(),
		}
	}
	return handlers.IngredientPackageHandlerInstance
}

func GetIngredientPackageServiceInstance() *services.IngredientPackageService {
	if services.IngredientPackageServiceInstance == nil {
		services.IngredientPackageServiceInstance = &services.IngredientPackageService{
			IngredientPackageDao: *GetIngredientPackageDaoInstance(),
			PackageDao:           *GetPackageDaoInstance(),
		}
	}
	return services.IngredientPackageServiceInstance
}

func GetIngredientPackageDaoInstance() *dao.IngredientPackageDao {
	if dao.IngredientPackageDaoInstance == nil {
		dao.IngredientPackageDaoInstance = &dao.IngredientPackageDao{
			IngredientCollection: core.GetDatabaseConnection().Collection("ingredients"),
		}
	}
	return dao.IngredientPackageDaoInstance
}
