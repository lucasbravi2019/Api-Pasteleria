package ingredients

import "github.com/lucasbravi2019/pasteleria/core"

func GetIngredientHandlerInstance() *ingredientHandler {
	if ingredientHandlerInstance == nil {
		ingredientHandlerInstance = &ingredientHandler{
			ingredientService: GetIngredientServiceInstance(),
		}
	}
	return ingredientHandlerInstance
}

func GetIngredientServiceInstance() *ingredientService {
	if ingredientServiceInstance == nil {
		ingredientServiceInstance = &ingredientService{
			repository: GetIngredientRepositoryInstance(),
		}
	}
	return ingredientServiceInstance
}

func GetIngredientRepositoryInstance() *ingredientRepository {
	if ingredientRepositoryInstance == nil {
		ingredientRepositoryInstance = &ingredientRepository{
			db: core.GetDatabaseConnection().Collection("ingredients"),
		}
	}
	return ingredientRepositoryInstance
}
