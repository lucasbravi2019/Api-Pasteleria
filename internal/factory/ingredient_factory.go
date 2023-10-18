package factory

import (
	"github.com/lucasbravi2019/pasteleria/api"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dao"
	"github.com/lucasbravi2019/pasteleria/internal/services"
)

func GetIngredientHandlerInstance() *api.IngredientHandler {
	if api.IngredientHandlerInstance == nil {
		api.IngredientHandlerInstance = &api.IngredientHandler{
			Service: GetIngredientServiceInstance(),
		}
	}
	return api.IngredientHandlerInstance
}

func GetIngredientServiceInstance() *services.IngredientService {
	if services.IngredientServiceInstance == nil {
		services.IngredientServiceInstance = &services.IngredientService{
			IngredientDao: *GetIngredientDaoInstance(),
		}
	}
	return services.IngredientServiceInstance
}

func GetIngredientDaoInstance() *dao.IngredientDao {
	if dao.IngredientDaoInstance == nil {
		dao.IngredientDaoInstance = &dao.IngredientDao{
			DB: db.GetDatabaseConnection(),
		}
	}
	return dao.IngredientDaoInstance
}
