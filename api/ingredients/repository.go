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
	CreateIngredient(ingredient *IngredientDTO) (int, *IngredientDTO)
	UpdateIngredient(oid *primitive.ObjectID, dto *IngredientDTO) (int, *IngredientDTO)
	DeleteIngredient(oid *primitive.ObjectID) (int, *IngredientDTO)
	AddPackageToIngredient(dto IngredientPackageDTO) (int, *IngredientDTO)
	ChangeIngredientPrice(packageOid *primitive.ObjectID, priceDTO *IngredientPackagePriceDTO) (int, *IngredientDTO)
}

var ingredientRepositoryInstance *repository

func (r *repository) GetAllIngredients() (int, []IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results, err := r.ingredientCollection.Aggregate(ctx, GetAggregateShowIngredients())

	if err != nil {
		log.Println(err.Error())
	}

	var ingredients []IngredientDTO

	err = results.All(ctx, &ingredients)

	if err != nil {
		log.Println(err.Error())
	}

	if len(ingredients) < 1 {
		return http.StatusOK, []IngredientDTO{}
	}

	return http.StatusOK, ingredients
}

func (r *repository) FindIngredientByOID(oid *primitive.ObjectID) (int, *IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var ingredient []IngredientDTO = []IngredientDTO{}

	cursor, err := r.ingredientCollection.Aggregate(ctx, GetIngredientById(*oid))

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	err = cursor.All(ctx, &ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	if len(ingredient) == 0 {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, &ingredient[0]
}

func (r *repository) CreateIngredient(ingredient *IngredientDTO) (int, *IngredientDTO) {
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

func (r *repository) UpdateIngredient(oid *primitive.ObjectID, dto *IngredientDTO) (int, *IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.ingredientCollection.UpdateOne(ctx, GetIngredientById(*oid), UpdateIngredientName(*dto))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, dto
}

func (r *repository) DeleteIngredient(oid *primitive.ObjectID) (int, *IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var ingredientDeleted *IngredientDTO = &IngredientDTO{}
	err := r.ingredientCollection.FindOneAndDelete(ctx, GetIngredientById(*oid)).Decode(ingredientDeleted)

	if err != nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, ingredientDeleted
}

func (r *repository) AddPackageToIngredient(dto IngredientPackageDTO) (int, *IngredientDTO) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	var ingredientPackage *IngredientPackage = &IngredientPackage{
		ID:    dto.PackageOid,
		Price: dto.Price,
	}

	_, err := r.ingredientCollection.UpdateOne(ctx, GetIngredientWithoutExistingPackage(dto.IngredientOid, dto.PackageOid),
		PushPackageIntoIngredient(*ingredientPackage))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, nil
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
