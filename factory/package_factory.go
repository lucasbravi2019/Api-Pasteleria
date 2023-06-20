package factory

import (
	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dao"
	"github.com/lucasbravi2019/pasteleria/handlers"
	"github.com/lucasbravi2019/pasteleria/services"
)

func GetPackageHandlerInstance() *handlers.PackageHandler {
	if handlers.PackageHandlerInstance == nil {
		handlers.PackageHandlerInstance = &handlers.PackageHandler{
			Service: GetPackageServiceInstance(),
		}
	}
	return handlers.PackageHandlerInstance
}

func GetPackageServiceInstance() *services.PackageService {
	if services.PackageServiceInstance == nil {
		services.PackageServiceInstance = &services.PackageService{
			PackageDao:           *GetPackageDaoInstance(),
			IngredientPackageDao: *GetIngredientPackageDaoInstance(),
			RecipeDao:            *GetRecipeDaoInstance(),
			RecipeIngredientDao:  *GetRecipeIngredientDaoInstance(),
		}
	}
	return services.PackageServiceInstance
}

func GetPackageDaoInstance() *dao.PackageDao {
	if dao.PackageDaoInstance == nil {
		dao.PackageDaoInstance = &dao.PackageDao{
			DB: core.GetDatabaseConnection().Collection("packages"),
		}
	}
	return dao.PackageDaoInstance
}
