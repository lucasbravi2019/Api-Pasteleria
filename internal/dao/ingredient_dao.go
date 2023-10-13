package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/mapper"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientDao struct {
	DB *sql.DB
}

type IngredientDaoInterface interface {
	GetAllIngredients() *[]dto.IngredientDTO
	FindIngredientByOID(oid *primitive.ObjectID) *dto.IngredientDTO
	FindIngredientByPackageId(packageId *primitive.ObjectID) *dto.IngredientDTO
	ValidateExistingIngredient(ingredientName *dto.IngredientNameDTO) error
	CreateIngredient(ingredient *models.Ingredient) (*primitive.ObjectID, error)
	UpdateIngredient(oid *primitive.ObjectID, dto *dto.IngredientNameDTO) error
	DeleteIngredient(oid *primitive.ObjectID) error
	ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *dto.IngredientPackagePriceDTO) error
}

var IngredientDaoInstance *IngredientDao

func (d *IngredientDao) GetAllIngredients() *[]dto.IngredientDTO {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	rows, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return mapper.ToIngredientDTOList(rows)
}

func (d *IngredientDao) FindIngredientByOID(oid *primitive.ObjectID) *dto.IngredientDTO {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	rows, err := d.DB.Query("")

	if err != nil {
		return nil
	}

	return mapper.ToIngredientDTO(rows)
}

func (d *IngredientDao) FindIngredientByPackageId(packageId *primitive.ObjectID) *dto.IngredientDTO {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	rows, err := d.DB.Query("")

	if err != nil {
		return nil
	}

	return mapper.ToIngredientDTO(rows)
}

func (d *IngredientDao) CreateIngredient(ingredient *models.Ingredient) {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}
}

func (d *IngredientDao) UpdateIngredient(oid *primitive.ObjectID, dto *dto.IngredientNameDTO) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		return err
	}

	return err
}

func (d *IngredientDao) DeleteIngredient(oid *primitive.ObjectID) error {
	_, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *IngredientDao) ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *dto.IngredientPackagePriceDTO) error {

	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}
