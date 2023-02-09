package ingredients

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/lucasbravi2019/pasteleria/api/packages"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type repository struct {
	ingredientCollection *mongo.Collection
	packageCollection    *mongo.Collection
	recipeCollection     *mongo.Collection
}

type IngredientRepository interface {
	GetAllIngredients() (int, []Ingredient)
	FindIngredientByOID(oid *primitive.ObjectID) (int, *Ingredient)
	CreateIngredient(ingredient *Ingredient) (int, *Ingredient)
	UpdateIngredient(oid *primitive.ObjectID, ingredient *Ingredient) (int, *Ingredient)
	DeleteIngredient(oid *primitive.ObjectID) (int, *Ingredient)
	AddPackageToIngredient(dto IngredientPackageDTO) (int, *Ingredient)
	ChangeIngredientPrice(*primitive.ObjectID, *IngredientPackagePrice) (int, *Ingredient)
}

var ingredientRepositoryInstance *repository

func (r *repository) GetAllIngredients() (int, []Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	results, err := r.ingredientCollection.Find(ctx, bson.D{})

	if err != nil {
		log.Println(err.Error())
	}

	var ingredients []Ingredient

	err = results.All(ctx, &ingredients)

	if err != nil {
		log.Println(err.Error())
	}

	if len(ingredients) < 1 {
		return http.StatusOK, []Ingredient{}
	}

	return http.StatusOK, ingredients
}

func (r *repository) FindIngredientByOID(oid *primitive.ObjectID) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var ingredient *Ingredient = &Ingredient{}

	err := r.ingredientCollection.FindOne(ctx, GetIngredientById(*oid)).Decode(ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, ingredient
}

func (r *repository) CreateIngredient(ingredient *Ingredient) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	pipelines := GetAggregateCreateIngredients(ingredient)

	cursor, err := r.ingredientCollection.Aggregate(ctx, pipelines)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	var ingredientsDuplicated *[]Ingredient = &[]Ingredient{}

	err = cursor.All(ctx, ingredientsDuplicated)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	if len(*ingredientsDuplicated) > 0 {
		return http.StatusBadRequest, nil
	}

	insertResult, err := r.ingredientCollection.InsertOne(ctx, *ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	oid := insertResult.InsertedID

	if oid == nil {
		return http.StatusInternalServerError, nil
	}

	var ingredientCreated *Ingredient = &Ingredient{
		ID:   oid.(primitive.ObjectID),
		Name: ingredient.Name,
	}

	return http.StatusCreated, ingredientCreated
}

func (r *repository) UpdateIngredient(oid *primitive.ObjectID, ingredient *Ingredient) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}
	_, err := r.ingredientCollection.ReplaceOne(ctx, filter, ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, ingredient
}

func (r *repository) DeleteIngredient(oid *primitive.ObjectID) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var ingredientDeleted *Ingredient = &Ingredient{}
	err := r.ingredientCollection.FindOneAndDelete(ctx, GetIngredientById(*oid)).Decode(ingredientDeleted)

	if err != nil {
		return http.StatusNotFound, nil
	}

	return http.StatusOK, ingredientDeleted
}

func (r *repository) AddPackageToIngredient(dto IngredientPackageDTO) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	var packageFound *packages.Package = &packages.Package{}

	err := r.packageCollection.FindOne(ctx, packages.GetPackageById(dto.PackageOid)).Decode(packageFound)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	packageFound.Price = dto.Price

	var envase *packages.Package = &packages.Package{
		ID:       packageFound.ID,
		Metric:   packageFound.Metric,
		Quantity: packageFound.Quantity,
		Price:    packageFound.Price,
	}

	var ingredient *Ingredient = &Ingredient{}

	err = r.ingredientCollection.FindOneAndUpdate(ctx, GetIngredientWithoutExistingPackage(dto.IngredientOid, dto.PackageOid),
		PushPackageIntoIngredient(*envase)).
		Decode(ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, ingredient
}

func (r *repository) ChangeIngredientPrice(ingredientPackageOid *primitive.ObjectID,
	ingredientPackagePrice *IngredientPackagePrice) (int, *Ingredient) {

	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	var envase *packages.Package = &packages.Package{}

	err := r.packageCollection.FindOne(ctx, packages.GetPackageById(*ingredientPackageOid)).Decode(envase)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	envase.Price = ingredientPackagePrice.Price

	var ingredient *Ingredient = &Ingredient{}

	_, err = r.ingredientCollection.UpdateOne(ctx, GetIngredientByPackageId(*ingredientPackageOid), SetIngredientPackages(*envase))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	err = r.ingredientCollection.FindOne(ctx, GetIngredientByPackageId(*ingredientPackageOid)).Decode(ingredient)

	if err != nil {
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, ingredient
}
