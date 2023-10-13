package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/core"
	"github.com/lucasbravi2019/pasteleria/dto"
<<<<<<< HEAD
	"github.com/lucasbravi2019/pasteleria/mapper"
=======
	"github.com/lucasbravi2019/pasteleria/models"
	"github.com/lucasbravi2019/pasteleria/queries"
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RecipeDao struct {
	DB *sql.DB
}

type RecipeDaoInterface interface {
	FindAllRecipes() *[]dto.RecipeDTO
	FindRecipeByOID(oid *primitive.ObjectID) (*dto.RecipeDTO, error)
	FindRecipesByPackageId(oid *primitive.ObjectID) ([]dto.RecipeDTO, error)
	CreateRecipe(recipe *models.Recipe) (*primitive.ObjectID, error)
	UpdateRecipeName(oid *primitive.ObjectID, recipeName *dto.RecipeNameDTO) error
	DeleteRecipe(oid *primitive.ObjectID) error
	UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error
	UpdateRecipesPrice() error
	GetRecipesByIngredientId(oid *primitive.ObjectID) (*[]models.Recipe, error)
	UpdateRecipes(recipes *[]models.Recipe) error
}

var RecipeDaoInstance *RecipeDao

<<<<<<< HEAD
func (d *RecipeDao) FindAllRecipes() *[]dto.RecipeDTO {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
=======
func (d *RecipeDao) FindAllRecipes() (*[]dto.RecipeDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	defer cancel()
	sql, err := core.FindQueryByName(core.Recipe_FindAll)

	if err != nil {
<<<<<<< HEAD
		log.Println(err)
		return nil
=======
		log.Println(err.Error())
		return recipes, err
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	}

	rows, err := d.DB.Query(sql)

<<<<<<< HEAD
	defer rows.Close()

	return mapper.ToRecipeDTOList(rows)
}

func (d *RecipeDao) FindRecipeByOID(oid *primitive.ObjectID) *dto.RecipeDTO {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
=======
	if err != nil {
		log.Println(err.Error())
		return recipes, err
	}

	return recipes, nil
}

func (d *RecipeDao) FindRecipeByOID(oid *primitive.ObjectID) (*dto.RecipeDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	defer cancel()

	rows, err := d.DB.Query("")

	if err != nil {
<<<<<<< HEAD
		return nil
	}

	return mapper.ToRecipeDTO(rows)
}

func (d *RecipeDao) FindRecipesByPackageId(packageId *primitive.ObjectID) []dto.RecipeDTO {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
=======
		log.Println(err.Error())
		return nil, err
	}

	return recipe, nil
}

func (d *RecipeDao) FindRecipesByPackageId(packageId *primitive.ObjectID) ([]dto.RecipeDTO, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	defer cancel()

	rows, err := d.DB.Query("")

	if err != nil {
<<<<<<< HEAD
		return nil
	}

	return *mapper.ToRecipeDTOList(rows)
}

func (d *RecipeDao) CreateRecipe(recipe *dto.RecipeNameDTO) error {
=======
		log.Println(err.Error())
		return nil, err
	}

	err = cursor.All(ctx, recipes)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return *recipes, nil
}

func (d *RecipeDao) CreateRecipe(recipe *models.Recipe) (*primitive.ObjectID, error) {
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	tx, err := d.DB.Begin()

	if err != nil {
<<<<<<< HEAD
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
=======
		log.Println(err.Error())
		return nil, err
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id, nil
}

func (d *RecipeDao) UpdateRecipeName(oid *primitive.ObjectID, recipe *models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.UpdateOne(ctx, queries.GetRecipeById(*oid), queries.UpdateRecipeName(*recipe))
>>>>>>> 9e63822ae2f7c13e69bf82f4c317e471e2a1be2e

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

func (d *RecipeDao) GetRecipesByIngredientId(oid *primitive.ObjectID) (*[]models.Recipe, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cur, err := d.DB.Find(ctx, queries.GetRecipeByIngredientId(*oid))

	if err != nil {
		return nil, err
	}

	recipes := &[]models.Recipe{}

	err = cur.All(ctx, recipes)

	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (d *RecipeDao) UpdateRecipes(recipes []models.Recipe) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	for i := 0; i < len(recipes); i++ {
		_, err := d.DB.UpdateOne(ctx, queries.GetRecipeById(recipes[i].ID), queries.UpdateRecipe(recipes[i]))

		if err != nil {
			return err
		}
	}

	return nil
}
