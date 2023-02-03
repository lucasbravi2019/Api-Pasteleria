package packages

import "github.com/lucasbravi2019/pasteleria/core"

func GetPackageHandlerInstance() *handler {
	if packageHandlerInstance == nil {
		packageHandlerInstance = &handler{
			service: GetPackageServiceInstance(),
		}
	}
	return packageHandlerInstance
}

func GetPackageServiceInstance() *service {
	if packageServiceInstance == nil {
		packageServiceInstance = &service{
			repository: GetPackageRepositoryInstance(),
		}
	}
	return packageServiceInstance
}

func GetPackageRepositoryInstance() *repository {
	if packageRepositoryInstance == nil {
		packageRepositoryInstance = &repository{
			db: core.GetDatabaseConnection().Collection("packages"),
		}
	}
	return packageRepositoryInstance
}
