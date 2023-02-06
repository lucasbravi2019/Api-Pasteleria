package ingredients

import (
	"context"
	"log"
	"net/http"
	"strings"
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
	ChangeIngredientPrice(*primitive.ObjectID, *IngredientPackagePrice) (int, *[]Ingredient)
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

	filter := bson.M{"_id": oid}

	result := r.ingredientCollection.FindOne(ctx, filter)

	if result.Err() != nil {
		log.Println(result.Err())
		return http.StatusNotFound, nil
	}

	var ingredient *Ingredient = &Ingredient{}

	err := result.Decode(ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, ingredient
}

func (r *repository) CreateIngredient(ingredient *Ingredient) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	project := bson.D{
		{Key: "$project", Value: bson.D{
			{Key: "name", Value: bson.D{
				{Key: "$toLower", Value: "$name"},
			}},
		}},
	}

	match := bson.D{
		{Key: "$match", Value: bson.D{
			{Key: "name", Value: strings.ToLower(ingredient.Name)},
		}},
	}

	cursor, err := r.ingredientCollection.Aggregate(ctx, mongo.Pipeline{project, match})

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
	result, err := r.ingredientCollection.ReplaceOne(ctx, filter, ingredient)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	uid := result.UpsertedID

	if uid == nil {
		return http.StatusInternalServerError, nil
	}

	var ingredientUpdated *Ingredient = &Ingredient{
		ID:       uid.(primitive.ObjectID),
		Name:     ingredient.Name,
		Packages: ingredient.Packages,
	}

	return http.StatusOK, ingredientUpdated
}

func (r *repository) DeleteIngredient(oid *primitive.ObjectID) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filter := bson.M{"_id": oid}

	result := r.ingredientCollection.FindOneAndDelete(ctx, filter)

	var ingredientDeleted *Ingredient = &Ingredient{}

	err := result.Decode(ingredientDeleted)

	if err != nil {
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, ingredientDeleted
}

func (r *repository) AddPackageToIngredient(dto IngredientPackageDTO) (int, *Ingredient) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	ingredientFilter := bson.M{"_id": dto.IngredientOid}

	result := r.ingredientCollection.FindOne(ctx, ingredientFilter)

	var ingredientFound *Ingredient = &Ingredient{}

	err := result.Decode(ingredientFound)

	if err != nil {
		log.Println("Ingredient not found")
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	packageFilter := bson.M{"_id": dto.PackageOid}

	result = r.packageCollection.FindOne(ctx, packageFilter)

	var packageFound *packages.Package = &packages.Package{}

	err = result.Decode(packageFound)

	if err != nil {
		log.Println("Package not found")
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	var ingredientPackage *IngredientPackage = &IngredientPackage{
		Package: *packageFound,
		Price:   dto.Price,
	}

	var anotherPackageExists bool = false

	for i := 0; i < len(ingredientFound.Packages); i++ {
		if ingredientFound.Packages[i].Package.ID == ingredientPackage.Package.ID {
			ingredientFound.Packages[i] = *ingredientPackage
			anotherPackageExists = true
			break
		}
	}

	if !anotherPackageExists {
		ingredientFound.Packages = append(ingredientFound.Packages, *ingredientPackage)
	}

	document := bson.M{"$set": bson.M{
		"packages": ingredientFound.Packages,
	}}

	updateResult, err := r.ingredientCollection.UpdateByID(ctx, dto.IngredientOid, document)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	if updateResult.ModifiedCount < 1 {
		log.Println("El ingrediente no fue actualizado")
	}

	return http.StatusOK, ingredientFound
}

func (r *repository) ChangeIngredientPrice(ingredientPackageOid *primitive.ObjectID, ingredientPackagePrice *IngredientPackagePrice) (int, *[]Ingredient) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	filter := bson.M{"packages.package._id": ingredientPackageOid}

	cursor, err := r.ingredientCollection.Find(ctx, filter)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	var ingredients []Ingredient

	err = cursor.All(ctx, &ingredients)

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	for i := 0; i < len(ingredients); i++ {
		for j := 0; j < len(ingredients[i].Packages); j++ {
			envase := ingredients[i].Packages[j]
			if envase.Package.ID == *ingredientPackageOid {
				ingredients[i].Packages[j].Price = ingredientPackagePrice.Price
				document := bson.M{"$set": bson.M{
					"packages": ingredients[i].Packages,
				}}
				updateResult, err := r.ingredientCollection.UpdateByID(ctx, ingredients[i].ID, document)

				if err != nil {
					log.Println(err.Error())
					return http.StatusInternalServerError, nil
				}

				if updateResult.MatchedCount == 0 {
					return http.StatusNotFound, nil
				}
			}
		}
	}

	return http.StatusOK, &ingredients
}
