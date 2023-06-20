package ingredients

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	ingredientCollection *mongo.Collection
	packageCollection    *mongo.Collection
	recipeCollection     *mongo.Collection
}

type IngredientRepository interface {
	GetAllIngredients() []IngredientDTO
	FindIngredientByOID(oid *primitive.ObjectID) *IngredientDTO
	FindIngredientByPackageId(packageId *primitive.ObjectID) *IngredientDTO
	ValidateExistingIngredient(ingredientName *IngredientNameDTO) error
	CreateIngredient(ingredient *Ingredient) *primitive.ObjectID
	UpdateIngredient(oid *primitive.ObjectID, dto *IngredientNameDTO) error
	DeleteIngredient(oid *primitive.ObjectID) error
	AddPackageToIngredient(ingredientOid *primitive.ObjectID, packageOid *primitive.ObjectID, envase *IngredientPackage) error
	RemovePackageFromIngredients(dto IngredientPackageDTO) error
	ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *IngredientPackagePriceDTO) error
}

var ingredientRepositoryInstance *repository

func (r *repository) GetAllIngredients() []IngredientDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results, err := r.ingredientCollection.Find(ctx, All())

	if err != nil {
		log.Println(err.Error())
	}

	var ingredients *[]IngredientDTO = &[]IngredientDTO{}

	err = results.All(ctx, ingredients)

	if err != nil {
		log.Println(err.Error())
	}

	if len(*ingredients) < 1 {
		return []IngredientDTO{}
	}

	return *ingredients
}

func (r *repository) FindIngredientByOID(oid *primitive.ObjectID) *IngredientDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var ingredient *IngredientDTO = &IngredientDTO{}

	err := r.ingredientCollection.FindOne(ctx, GetIngredientById(*oid)).Decode(ingredient)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return ingredient
}

func (r *repository) FindIngredientByPackageId(packageId *primitive.ObjectID) *IngredientDTO {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	var ingredient *IngredientDTO = &IngredientDTO{}

	err := r.ingredientCollection.FindOne(ctx, GetIngredientByPackageId(*packageId)).Decode(ingredient)

	if err != nil {
		return nil
	}

	return ingredient
}

func (r *repository) CreateIngredient(ingredient *Ingredient) *primitive.ObjectID {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	insertResult, err := r.ingredientCollection.InsertOne(ctx, *ingredient)

	if err != nil {
		log.Println(err.Error())
	}

	id := insertResult.InsertedID.(primitive.ObjectID)

	return &id
}

func (r *repository) UpdateIngredient(oid *primitive.ObjectID, dto *IngredientNameDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.UpdateOne(ctx, GetIngredientById(*oid), UpdateIngredientName(*dto))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *repository) DeleteIngredient(oid *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.DeleteOne(ctx, GetIngredientById(*oid))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *repository) AddPackageToIngredient(ingredientOid *primitive.ObjectID, packageOid *primitive.ObjectID, envase *IngredientPackage) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.UpdateOne(ctx, GetIngredientWithoutExistingPackage(*ingredientOid, *packageOid), PushPackageIntoIngredient(*envase))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *repository) RemovePackageFromIngredients(dto IngredientPackageDTO) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.UpdateMany(ctx, GetIngredientByPackageId(dto.PackageOid), PullPackageFromIngredients(dto))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *repository) ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *IngredientPackagePriceDTO) error {

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.UpdateOne(ctx, GetIngredientByPackageId(*packageOid), SetIngredientPrice(priceDTO.Price), GetArrayFilterForPackageId(*packageOid))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (r *repository) ValidateExistingIngredient(ingredientName *IngredientNameDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := r.ingredientCollection.Aggregate(ctx, GetAggregateCreateIngredients(ingredientName))

	if err != nil {
		log.Println(err.Error())
		return err
	}

	var ingredientsDuplicated *[]IngredientDTO = &[]IngredientDTO{}

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
