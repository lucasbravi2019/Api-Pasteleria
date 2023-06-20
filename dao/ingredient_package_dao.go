package dao

import (
	"context"
	"log"
	"time"

	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"github.com/lucasbravi2019/pasteleria/queries"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IngredientPackageDao struct {
	IngredientCollection *mongo.Collection
}

type IngredientPackageDaoInterface interface {
	UpdateIngredientPackagePrice(packageId *primitive.ObjectID, price float64) error
	AddPackageToIngredient(ingredientOid *primitive.ObjectID, packageOid *primitive.ObjectID, envase *models.IngredientPackage) error
}

var IngredientPackageDaoInstance *IngredientPackageDao

func (d *IngredientPackageDao) AddPackageToIngredient(ingredientOid *primitive.ObjectID, packageOid *primitive.ObjectID, envase *models.IngredientPackage) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.IngredientCollection.UpdateOne(ctx, queries.GetIngredientWithoutExistingPackage(*ingredientOid, *packageOid),
		queries.PushPackageIntoIngredient(*envase))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *IngredientPackageDao) RemovePackageFromIngredients(dto dto.IngredientPackageDTO) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.IngredientCollection.UpdateMany(ctx, queries.GetIngredientByPackageId(dto.PackageOid), queries.PullPackageFromIngredients(dto))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
