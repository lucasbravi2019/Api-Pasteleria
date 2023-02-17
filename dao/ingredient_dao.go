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

type IngredientDao struct {
	IngredientCollection *mongo.Collection
}

type IngredientDaoInterface interface {
	GetAllIngredients() []dto.IngredientDTO
	FindIngredientByOID(oid *primitive.ObjectID) *dto.IngredientDTO
	FindIngredientByPackageId(packageId *primitive.ObjectID) *dto.IngredientDTO
	ValidateExistingIngredient(ingredientName *dto.IngredientNameDTO) error
	CreateIngredient(ingredient *models.Ingredient) *primitive.ObjectID
	UpdateIngredient(oid *primitive.ObjectID, dto *dto.IngredientNameDTO) error
	DeleteIngredient(oid *primitive.ObjectID) error
	ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *dto.IngredientPackagePriceDTO) error
}

var IngredientDaoInstance *IngredientDao

func (d *IngredientDao) GetAllIngredients() []dto.IngredientDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results, err := d.IngredientCollection.Find(ctx, queries.All())

	if err != nil {
		log.Println(err.Error())
	}

	ingredients := &[]dto.IngredientDTO{}

	err = results.All(ctx, ingredients)

	if err != nil {
		log.Println(err.Error())
	}

	if len(*ingredients) < 1 {
		return []dto.IngredientDTO{}
	}

	return *ingredients
}

func (d *IngredientDao) FindIngredientByOID(oid *primitive.ObjectID) *dto.IngredientDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	ingredient := &dto.IngredientDTO{}

	err := d.IngredientCollection.FindOne(ctx, queries.GetIngredientById(*oid)).Decode(ingredient)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return ingredient
}

func (d *IngredientDao) FindIngredientByPackageId(packageId *primitive.ObjectID) *dto.IngredientDTO {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	ingredient := &dto.IngredientDTO{}

	err := d.IngredientCollection.FindOne(ctx, queries.GetIngredientByPackageId(*packageId)).Decode(ingredient)

	if err != nil {
		return nil
	}

	return ingredient
}

func (d *IngredientDao) CreateIngredient(ingredient *models.Ingredient) *primitive.ObjectID {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	insertResult, err := d.IngredientCollection.InsertOne(ctx, *ingredient)

	if err != nil {
		log.Println(err.Error())
	}

	id := insertResult.InsertedID.(primitive.ObjectID)

	return &id
}

func (d *IngredientDao) UpdateIngredient(oid *primitive.ObjectID, dto *dto.IngredientNameDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.IngredientCollection.UpdateOne(ctx, queries.GetIngredientById(*oid), queries.UpdateIngredientName(*dto))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *IngredientDao) DeleteIngredient(oid *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := d.IngredientCollection.DeleteOne(ctx, queries.GetIngredientById(*oid))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (d *IngredientDao) ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *dto.IngredientPackagePriceDTO) error {

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := d.IngredientCollection.UpdateOne(ctx, queries.GetIngredientByPackageId(*packageOid), queries.SetIngredientPrice(priceDTO.Price),
		queries.GetArrayFilterForPackageId(*packageOid))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (d *IngredientDao) ValidateExistingIngredient(ingredientName *dto.IngredientNameDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := d.IngredientCollection.Aggregate(ctx, queries.GetAggregateCreateIngredients(ingredientName))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	ingredientsDuplicated := &[]dto.IngredientDTO{}

	err = cursor.All(ctx, ingredientsDuplicated)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	if len(*ingredientsDuplicated) > 0 {
		return err
	}

	return nil
}
