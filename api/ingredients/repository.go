package ingredients

import (
	"context"
	"log"
	"net/http"
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
	GetAllIngredients() (int, []IngredientDTO)
	FindIngredientByOID(oid *primitive.ObjectID) (int, *IngredientDTO)
	CreateIngredient(ingredient *IngredientNameDTO) (int, *IngredientDTO)
	UpdateIngredient(oid *primitive.ObjectID, dto *IngredientNameDTO) (int, *IngredientDTO)
	DeleteIngredient(oid *primitive.ObjectID) (int, *primitive.ObjectID)
	AddPackageToIngredient(ingredientOid *primitive.ObjectID, packageOid *primitive.ObjectID, envase *IngredientPackage) (int, *IngredientDTO)
	RemovePackageFromIngredients(dto IngredientPackageDTO) (int, *primitive.ObjectID)
	ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *IngredientPackagePriceDTO) (int, *IngredientDTO)
}

var ingredientRepositoryInstance *repository

func (r *repository) GetAllIngredients() (int, []IngredientDTO) {
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
		return http.StatusOK, []IngredientDTO{}
	}

	return http.StatusOK, *ingredients
}

func (r *repository) FindIngredientByOID(oid *primitive.ObjectID) (int, *IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var ingredient *IngredientDTO = &IngredientDTO{}

	err := r.ingredientCollection.FindOne(ctx, GetIngredientById(*oid)).Decode(ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, ingredient
}

func (r *repository) CreateIngredient(ingredient *IngredientNameDTO) (int, *IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := r.ingredientCollection.Aggregate(ctx, GetAggregateCreateIngredients(ingredient))

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	var ingredientsDuplicated *[]IngredientDTO = &[]IngredientDTO{}

	err = cursor.All(ctx, ingredientsDuplicated)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	if len(*ingredientsDuplicated) > 0 {
		return http.StatusBadRequest, nil
	}

	var ingredientEntity *Ingredient = &Ingredient{
		Name:     ingredient.Name,
		Packages: []IngredientPackage{},
	}

	insertResult, err := r.ingredientCollection.InsertOne(ctx, *ingredientEntity)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	oid := insertResult.InsertedID

	if oid == nil {
		return http.StatusInternalServerError, nil
	}

	var ingredientCreated *IngredientDTO = &IngredientDTO{
		ID:   oid.(primitive.ObjectID),
		Name: ingredient.Name,
	}

	return http.StatusCreated, ingredientCreated
}

func (r *repository) UpdateIngredient(oid *primitive.ObjectID, dto *IngredientNameDTO) (int, *IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.UpdateOne(ctx, GetIngredientById(*oid), UpdateIngredientName(*dto))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	var ingredient *IngredientDTO = &IngredientDTO{}

	err = r.ingredientCollection.FindOne(ctx, GetIngredientById(*oid)).Decode(ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, ingredient
}

func (r *repository) DeleteIngredient(oid *primitive.ObjectID) (int, *primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.DeleteOne(ctx, GetIngredientById(*oid))

	if err != nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, oid
}

func (r *repository) AddPackageToIngredient(ingredientOid *primitive.ObjectID, packageOid *primitive.ObjectID, envase *IngredientPackage) (int, *IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.UpdateOne(ctx, GetIngredientWithoutExistingPackage(*ingredientOid, *packageOid), PushPackageIntoIngredient(*envase))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, nil
}

func (r *repository) RemovePackageFromIngredients(dto IngredientPackageDTO) (int, *primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.UpdateMany(ctx, GetIngredientByPackageId(dto.PackageOid), PullPackageFromIngredients(dto))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, &dto.PackageOid
}

func (r *repository) ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *IngredientPackagePriceDTO) (int, *IngredientDTO) {

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	var ingredient *IngredientDTO = &IngredientDTO{}

	_, err := r.ingredientCollection.UpdateOne(ctx, GetIngredientByPackageId(*packageOid), SetIngredientPrice(priceDTO.Price), GetArrayFilterForPackageId(*packageOid))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	err = r.ingredientCollection.FindOne(ctx, GetIngredientByPackageId(*packageOid)).Decode(ingredient)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, ingredient
}
