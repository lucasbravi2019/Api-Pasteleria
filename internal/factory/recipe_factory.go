package factory

import (
	"github.com/lucasbravi2019/pasteleria/api"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/services"
)

func GetRecipeHandlerInstance() *api.RecipeHandler {
	if api.RecipeHandlerInstance == nil {
		api.RecipeHandlerInstance = &api.RecipeHandler{
			Service: GetRecipeServiceInstance(),
		}
	}
	return api.RecipeHandlerInstance
}

func GetRecipeServiceInstance() *services.RecipeService {
	if services.RecipeServiceInstance == nil {
		services.RecipeServiceInstance = &services.RecipeService{
			RecipeDao:    GetRecipeDaoInstance(),
			RecipeMapper: GetRecipeMapperInstance(),
		}
	}
	return services.RecipeServiceInstance
}

func GetRecipeDaoInstance() *dao.RecipeDao {
	if dao.RecipeDaoInstance == nil {
		dao.RecipeDaoInstance = &dao.RecipeDao{
			DB:           db.GetDatabaseConnection(),
			RecipeMapper: GetRecipeMapperInstance(),
		}
	}
	return dao.RecipeDaoInstance
}

func GetRecipeMapperInstance() *mapper.RecipeMapper {
	if mapper.RecipeMapperInstance == nil {
		mapper.RecipeMapperInstance = &mapper.RecipeMapper{}
	}
	return mapper.RecipeMapperInstance
}
