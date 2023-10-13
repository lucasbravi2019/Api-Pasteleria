package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/mapper"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeDao struct {
	DB *sql.DB
}

type RecipeDaoInterface interface {
	FindAllRecipes() *[]dto.RecipeDTO
	FindRecipeByOID(oid *primitive.ObjectID) *dto.RecipeDTO
	FindRecipesByPackageId(oid *primitive.ObjectID) []dto.RecipeDTO
	CreateRecipe(recipe *dto.RecipeNameDTO) *primitive.ObjectID
	UpdateRecipeName(oid *primitive.ObjectID, recipeName *dto.RecipeNameDTO) error

	DeleteRecipe(oid *primitive.ObjectID) error

	UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error

	UpdateRecipesPrice() error
}

var RecipeDaoInstance *RecipeDao

func (d *RecipeDao) FindAllRecipes() *[]dto.RecipeDTO {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	sql, err := core.FindQueryByName(core.Recipe_FindAll)

	if err != nil {
		log.Println(err)
		return nil
	}

	rows, err := d.DB.Query(sql)

	defer rows.Close()

	return mapper.ToRecipeDTOList(rows)
}

func (d *RecipeDao) FindRecipeByOID(oid *primitive.ObjectID) *dto.RecipeDTO {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	rows, err := d.DB.Query("")

	if err != nil {
		return nil
	}

	return mapper.ToRecipeDTO(rows)
}

func (d *RecipeDao) FindRecipesByPackageId(packageId *primitive.ObjectID) []dto.RecipeDTO {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	rows, err := d.DB.Query("")

	if err != nil {
		return nil
	}

	return *mapper.ToRecipeDTOList(rows)
}

func (d *RecipeDao) CreateRecipe(recipe *dto.RecipeNameDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	tx, err := d.DB.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	query := core.GetQueryByName(core.Recipe_Create)

	log.Println(query)
	log.Println(recipe.Name)

	res, err := tx.ExecContext(ctx, "insert into recipe ('name') values (?)", recipe.Name)

	if err != nil {
		return err
	}

	err = tx.Commit()
	log.Println(res.LastInsertId())
	log.Println(res.RowsAffected())
	return err
}

func (d *RecipeDao) UpdateRecipeName(oid *primitive.ObjectID, recipeName *dto.RecipeNameDTO) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeDao) DeleteRecipe(oid *primitive.ObjectID) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeDao) UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *RecipeDao) UpdateRecipesPrice() error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
