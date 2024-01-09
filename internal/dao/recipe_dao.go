package dao

import (
	"context"
	"database/sql"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type RecipeDao struct {
	DB           *sql.DB
	RecipeMapper *mapper.RecipeMapper
}

var RecipeDaoInstance *RecipeDao

func (d *RecipeDao) FindAllRecipes() (*[]dto.Recipe, error) {
	query, err := db.GetQueryByName(db.Recipe_FindAll)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query)

	if pkg.HasError(err) {
		return nil, err
	}
	defer rows.Close()

	return d.RecipeMapper.ToRecipeList(rows)
}

func (d *RecipeDao) FindRecipeById(id int64) (*dto.Recipe, error) {
	query, err := db.GetQueryByName(db.Recipe_FindById)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query, id)

	if pkg.HasError(err) {
		return nil, err
	}

	return d.RecipeMapper.ToRecipeRow(rows)
}

func (d *RecipeDao) CreateRecipe(recipe *dto.RecipeRequest) error {
	query, err := db.GetQueryByName(db.Recipe_Create)

	if pkg.HasError(err) {
		return err
	}
	_, err = d.DB.Exec(query, recipe.Name)

	return err
}

func (d *RecipeDao) UpdateRecipe(recipe *dto.RecipeRequest) error {
	query, err := db.GetQueryByName(db.Recipe_UpdateName)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, recipe.Name, recipe.Id)

	return err
}

func (d *RecipeDao) DeleteRecipe(id *int64) error {
	query, err := db.GetQueryByName(db.Recipe_DeleteIngredientsByRecipeId)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, *id)

	if pkg.HasError(err) {
		return err
	}

	query, err = db.GetQueryByName(db.Recipe_DeleteById)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, *id)

	if pkg.HasError(err) {
		return err
	}

	return nil
}

func (d *RecipeDao) RemoveRecipeIngredientsByRecipeId(recipeId *int64) error {
	query, err := db.GetQueryByName(db.Recipe_DeleteIngredientsByRecipeId)

	if pkg.HasError(err) {
		return err
	}

	tx, err := d.DB.BeginTx(context.TODO(), nil)

	if pkg.HasError(err) {
		return err
	}

	defer func() {
		tx.Commit()
	}()

	_, err = d.DB.Exec(query, recipeId)

	if pkg.HasError(err) {
		return err
	}

	return nil
}

func (d *RecipeDao) AddRecipeIngredient(recipeId *int64, ingredients *[]dto.RecipeIngredientRequest) error {
	query, err := db.GetQueryByName(db.Recipe_AddIngredientsToRecipe)

	if pkg.HasError(err) {
		return err
	}

	tx, err := d.DB.BeginTx(context.TODO(), nil)

	if pkg.HasError(err) {
		return err
	}

	defer func() {
		tx.Commit()
	}()

	for _, ingredient := range *ingredients {
		_, err := d.DB.Exec(query, recipeId, ingredient.Id, ingredient.Quantity)

		if pkg.HasError(err) {
			return err
		}
	}

	return nil
}

func (d *RecipeDao) FindRecipeIdByName(recipeName *string) (*int64, error) {
	query, err := db.GetQueryByName(db.Recipe_FindRecipeIdByName)

	if pkg.HasError(err) {
		return nil, err
	}

	row := d.DB.QueryRow(query, recipeName)

	var recipeId int64

	err = row.Scan(&recipeId)

	if pkg.HasError(err) {
		return nil, err
	}

	return &recipeId, nil
}

func (d *RecipeDao) UpdateRecipePriceByRecipeId(recipeId *int64, price *float64) error {
	query, err := db.GetQueryByName(db.Recipe_UpdateRecipePriceByRecipeId)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, price, recipeId)

	return err
}

func (d *RecipeDao) UpdateRecipeIngredientPriceById(recipeIngredientId *int64, price *float64) error {
	query, err := db.GetQueryByName(db.Recipe_UpdateRecipeIngredientPriceById)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, price, recipeIngredientId)

	return err
}
