package dao

import (
	"database/sql"
	"errors"
	"log"
	"strconv"

	"github.com/lucasbravi2019/pasteleria/db"
	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeDao struct {
	DB *sql.DB
}

type RecipeDaoInterface interface {
	FindAllRecipes() *[]dto.RecipeDTO
	FindRecipeById(id int64) (*models.Recipe, error)
	FindRecipesByPackageId(oid *primitive.ObjectID) ([]dto.RecipeDTO, error)
	CreateRecipe(recipe *dto.RecipeNameDTO) error
	UpdateRecipeName(oid *primitive.ObjectID, recipeName *dto.RecipeNameDTO) error
	DeleteRecipe(id *int64) error
	UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error
	UpdateRecipesPrice() error
	GetRecipesByIngredientId(oid *primitive.ObjectID) (*[]models.Recipe, error)
	UpdateRecipes(recipes *[]models.Recipe) error
}

var RecipeDaoInstance *RecipeDao

func (d *RecipeDao) FindAllRecipes() (*[]models.Recipe, error) {
	query, err := db.GetQueryByName(db.Recipe_FindAll)

	if err != nil {
		return nil, err
	}
	log.Println(query)
	rows, err := d.DB.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	return mapper.ToRecipeList(rows), nil
}

func (d *RecipeDao) FindRecipeById(id int64) (*models.Recipe, error) {
	query, err := db.GetQueryByName(db.Recipe_FindById)

	if err != nil {
		return nil, err
	}

	row := d.DB.QueryRow(query, id)

	return mapper.ToRecipe(row), nil
}

func (d *RecipeDao) FindRecipesByPackageId(packageId *primitive.ObjectID) ([]dto.RecipeDTO, error) {
	_, err := d.DB.Query("")

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (d *RecipeDao) CreateRecipe(recipe *dto.RecipeNameDTO) error {
	query, err := db.GetQueryByName(db.Recipe_Create)

	if err != nil {
		return err
	}
	_, err = d.DB.Exec(query, recipe.Name)

	return err
}

func (d *RecipeDao) UpdateRecipeName(id *int64, recipe *dto.RecipeNameDTO) error {
	query, err := db.GetQueryByName(db.Recipe_UpdateName)

	if err != nil {
		return err
	}

	_, err = d.DB.Exec(query, recipe.Name, id)

	return err
}

func (d *RecipeDao) DeleteRecipe(id *int64) error {
	query, err := db.GetQueryByName(db.Recipe_DeleteById)

	if err != nil {
		return err
	}

	res, _ := d.DB.Exec(query, id)

	rowsAffected, err := res.RowsAffected()

	if err != nil || rowsAffected == 0 {
		return errors.New("rows affected: " + strconv.FormatInt(rowsAffected, 10))
	}

	return nil
}

func (d *RecipeDao) UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error {
	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeDao) UpdateRecipesPrice() error {
	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeDao) GetRecipesByIngredientId(oid *primitive.ObjectID) (*[]models.Recipe, error) {
	_, err := d.DB.Query("")

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (d *RecipeDao) UpdateRecipes(recipes []models.Recipe) error {

	return nil
}
