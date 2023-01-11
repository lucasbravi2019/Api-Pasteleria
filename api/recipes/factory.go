package recipes

import (
	"github.com/lucasbravi2019/pasteleria/core"
)

func GetRecipeHandlerInstance() *recipeHandler {
	if recipeHandlerInstance == nil {
		recipeHandlerInstance = &recipeHandler{
			service: GetRecipeServiceInstance(),
		}
	}
	return recipeHandlerInstance
}

func GetRecipeServiceInstance() *recipeService {
	if recipeServiceInstance == nil {
		recipeServiceInstance = &recipeService{
			repository: GetRecipeRepositoryInstance(),
		}
	}
	return recipeServiceInstance
}

func GetRecipeRepositoryInstance() *recipeRepository {
	if recipeRepositoryInstance == nil {
		recipeRepositoryInstance = &recipeRepository{
			db: core.GetDatabaseConnection().Collection("recipes"),
		}
	}
	return recipeRepositoryInstance
}
