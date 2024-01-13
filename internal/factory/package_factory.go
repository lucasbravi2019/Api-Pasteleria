package factory

import (
	"github.com/lucasbravi2019/pasteleria/api"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/services"
)

func GetPackageHandlerInstance() *api.PackageHandler {
	if api.PackageHandlerInstance == nil {
		api.PackageHandlerInstance = &api.PackageHandler{
			Service: GetPackageServiceInstance(),
		}
	}
	return api.PackageHandlerInstance
}

func GetPackageServiceInstance() *services.PackageService {
	if services.PackageServiceInstance == nil {
		services.PackageServiceInstance = &services.PackageService{
			PackageDao:    GetPackageDaoInstance(),
			PackageMapper: GetPackageMapperInstance(),
			RecipeService: GetRecipeServiceInstance(),
		}
	}
	return services.PackageServiceInstance
}

func GetPackageDaoInstance() *dao.PackageDao {
	if dao.PackageDaoInstance == nil {
		dao.PackageDaoInstance = &dao.PackageDao{
			DB:            db.GetDatabaseConnection(),
			PackageMapper: GetPackageMapperInstance(),
		}
	}
	return dao.PackageDaoInstance
}

func GetPackageMapperInstance() *mapper.PackageMapper {
	if mapper.PackageMapperInstance == nil {
		mapper.PackageMapperInstance = &mapper.PackageMapper{}
	}
	return mapper.PackageMapperInstance
}
