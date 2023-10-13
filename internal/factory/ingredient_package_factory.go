package factory

import (
	"github.com/lucasbravi2019/pasteleria/api"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/services"
)

func GetIngredientPackageHandlerInstance() *api.IngredientPackageHandler {
	if api.IngredientPackageHandlerInstance == nil {
		api.IngredientPackageHandlerInstance = &api.IngredientPackageHandler{
			Service: *GetIngredientPackageServiceInstance(),
		}
	}
	return api.IngredientPackageHandlerInstance
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
			DB: db.GetDatabaseConnection(),
		}
	}
	return dao.IngredientPackageDaoInstance
}
