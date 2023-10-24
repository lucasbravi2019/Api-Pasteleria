package dao

import (
	"database/sql"
	"fmt"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"github.com/lucasbravi2019/pasteleria/pkg"
)

type RecipeDao struct {
	DB *sql.DB
}

type RecipeDaoInterface interface {
	FindAllRecipes() *[]dto.RecipeDTO
	FindRecipeById(id int64) (*[]models.Recipe, error)
	CreateRecipe(recipe *dto.RecipeNameDTO) error
	UpdateRecipeName(recipeName *dto.RecipeNameDTO) error
	DeleteRecipe(id *int64) error
}

var RecipeDaoInstance *RecipeDao

func (d *RecipeDao) FindAllRecipes() (*[]models.Recipe, error) {
	query, err := db.GetQueryByName(db.Recipe_FindAll)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query)
	defer rows.Close()

	if pkg.HasError(err) {
		return nil, err
	}

	return mapper.ToRecipeList(rows), nil
}

func (d *RecipeDao) FindRecipeById(id int64) (*[]models.Recipe, error) {
	query, err := db.GetQueryByName(db.Recipe_FindById)

	if pkg.HasError(err) {
		return nil, err
	}

	rows, err := d.DB.Query(query, id)

	if pkg.HasError(err) {
		return nil, err
	}

	return mapper.ToRecipeList(rows), nil
}

func (d *RecipeDao) CreateRecipe(recipe *dto.RecipeNameDTO) error {
	query, err := db.GetQueryByName(db.Recipe_Create)

	if pkg.HasError(err) {
		return err
	}
	_, err = d.DB.Exec(query, recipe.Name)

	return err
}

func (d *RecipeDao) UpdateRecipeName(recipe *dto.RecipeNameDTO) error {
	query, err := db.GetQueryByName(db.Recipe_UpdateName)

	if pkg.HasError(err) {
		return err
	}

	_, err = d.DB.Exec(query, recipe.Name, recipe.Id)

	return err
}

func (d *RecipeDao) DeleteRecipe(id *int64) error {
	query, err := db.GetQueryByName(db.Recipe_DeleteById)

	if pkg.HasError(err) {
		return err
	}

	res, _ := d.DB.Exec(query, id)

	rowsAffected, err := res.RowsAffected()

	if pkg.HasError(err) {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("rows affected: %d", rowsAffected)
	}

	return nil
}
