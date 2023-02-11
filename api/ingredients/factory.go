package ingredients

import (
	"github.com/lucasbravi2019/pasteleria/api/packages"
	"github.com/lucasbravi2019/pasteleria/core"
)

func GetIngredientHandlerInstance() *handler {
	if ingredientHandlerInstance == nil {
		ingredientHandlerInstance = &handler{
			service: GetIngredientServiceInstance(),
		}
	}
	return ingredientHandlerInstance
}

func GetIngredientServiceInstance() *service {
	if ingredientServiceInstance == nil {
		ingredientServiceInstance = &service{
			ingredientRepository: GetIngredientRepositoryInstance(),
			packageRepository:    packages.GetPackageRepositoryInstance(),
		}
	}
	return ingredientServiceInstance
}

func GetIngredientRepositoryInstance() *repository {
	if ingredientRepositoryInstance == nil {
		ingredientRepositoryInstance = &repository{
			ingredientCollection: core.GetDatabaseConnection().Collection("ingredients"),
		}
	}
	return ingredientRepositoryInstance
}
