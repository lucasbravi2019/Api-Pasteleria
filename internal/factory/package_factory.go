package factory

import (
	"github.com/lucasbravi2019/pasteleria/api"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
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
			PackageDao: *GetPackageDaoInstance(),
		}
	}
	return services.PackageServiceInstance
}

func GetPackageDaoInstance() *dao.PackageDao {
	if dao.PackageDaoInstance == nil {
		dao.PackageDaoInstance = &dao.PackageDao{
			DB: db.GetDatabaseConnection(),
		}
	}
	return dao.PackageDaoInstance
}
