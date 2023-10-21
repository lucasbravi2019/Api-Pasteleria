package dao

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type RecipeIngredientDao struct {
	DB *sql.DB
}

type RecipeIngredientDaoInterface interface {
	GetAllRecipeIngredients(recipeId int64) (*[]models.RecipeIngredient, error)
	UpdateRecipeIngredients(ctx *gin.Context) error
}

var RecipeIngredientDaoInstance *RecipeIngredientDao

func (d *RecipeIngredientDao) GetAllRecipeIngredients(recipeId int64) (*[]models.RecipeIngredient, error) {
	query, err := db.GetQueryByName(db.Recipe_FindRecipeIngredients)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query, recipeId)

	if pkg.HasError(err) {
		return nil, err
	}

	recipes, err := mapper.ToRecipeIngredientList(rows)

	if pkg.HasError(err) {
		return nil, err
	}

	return recipes, nil
}
