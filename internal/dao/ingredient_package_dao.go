package dao

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/internal/dto"
	"github.com/lucasbravi2019/pasteleria/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IngredientPackageDao struct {
	DB *sql.DB
}

type IngredientPackageDaoInterface interface {
	UpdateIngredientPackagePrice(packageId *primitive.ObjectID, price float64) error
	AddPackageToIngredient(ingredientOid *primitive.ObjectID, packageOid *primitive.ObjectID, envase *models.IngredientPackage) error
}

var IngredientPackageDaoInstance *IngredientPackageDao

func (d *IngredientPackageDao) AddPackageToIngredient(ingredientOid *primitive.ObjectID, packageOid *primitive.ObjectID, envase *models.IngredientPackage) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *IngredientPackageDao) RemovePackageFromIngredients(dto dto.IngredientPackageDTO) error {
	_, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.DB.Query("")

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
