package dao

import (
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type RecipeIngredientDao struct {
	DB *sql.DB
}

type RecipeIngredientDaoInterface interface {
	GetAllRecipeIngredients(recipeId int64) (*[]models.RecipeIngredient, error)
	UpdateRecipeIngredients(ingredients dto.RecipeIngredientIdDTO) error
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

func (d *RecipeIngredientDao) UpdateRecipeIngredients(ingredients dto.RecipeIngredientIdDTO) error {
	tx, err := d.DB.Begin()
	if pkg.HasError(err) {
		return err
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	query, err := db.GetQueryByName(db.Recipe_DeleteIngredientsByRecipeId)

	if pkg.HasError(err) {
		return err
	}

	_, err = tx.Exec(query, ingredients.RecipeId)

	if pkg.HasError(err) {
		return err
	}

	query, err = db.GetQueryByName(db.Recipe_AddIngredientsToRecipe)

	if err != nil {
		return err
	}

	for _, ingredient := range ingredients.Ingredients {
		_, err = tx.Exec(query, ingredients.RecipeId, ingredient.IngredientId, ingredient.Quantity)

		if pkg.HasError(err) {
			return err
		}
	}

	return nil
}
