package dao

import (
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

type RecipeDaoInterface interface {
	FindAllRecipes() *[]dto.RecipeDTO
	FindRecipeById(id int64) (*dto.RecipeDTO, error)
	CreateRecipe(recipe *dto.RecipeCreationDTO) error
	UpdateRecipe(recipeName *dto.RecipeUpdateDTO) error
	DeleteRecipe(id *int64) error
}

var RecipeDaoInstance *RecipeDao

func (d *RecipeDao) FindAllRecipes() (*[]dto.RecipeDTO, error) {
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

func (d *RecipeDao) FindRecipeById(id int64) (*dto.RecipeDTO, error) {
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

func (d *RecipeDao) CreateRecipe(recipe *dto.RecipeCreationDTO) error {
	query, err := db.GetQueryByName(db.Recipe_Create)

	if pkg.HasError(err) {
		return err
	}
	_, err = d.DB.Exec(query, recipe.Name)

	return err
}

func (d *RecipeDao) UpdateRecipe(recipe *dto.RecipeUpdateDTO) error {
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
