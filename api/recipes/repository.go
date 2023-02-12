package recipes

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type recipeRepository struct {
	db *mongo.Collection
}

type RecipeRepository interface {
	FindAllRecipes() (int, *[]RecipeDTO)
	FindRecipeByOID(oid *primitive.ObjectID) (int, *RecipeDTO)
	CreateRecipe(recipe *RecipeNameDTO) (int, *RecipeDTO)
	UpdateRecipeName(oid *primitive.ObjectID, recipeName *RecipeNameDTO) (int, *RecipeDTO)
	AddIngredientToRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) (int, *RecipeDTO)
	RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) (int, *RecipeDTO)
	DeleteRecipe(oid *primitive.ObjectID) (int, *primitive.ObjectID)
	RemoveIngredientByPackageId(packageId *primitive.ObjectID) (int, *primitive.ObjectID)
	UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) (int, *primitive.ObjectID)
	UpdateRecipesPrice() int
	UpdateIngredientsPrice(ingredientId *primitive.ObjectID, price float64) (int, *primitive.ObjectID)
}

var recipeRepositoryInstance *recipeRepository

func (r *recipeRepository) FindAllRecipes() (int, *[]RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := r.db.Find(ctx, All())

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	var recipes *[]RecipeDTO = &[]RecipeDTO{}
	err = cursor.All(ctx, recipes)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipes
}

func (r *recipeRepository) FindRecipeByOID(oid *primitive.ObjectID) (int, *RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var recipe *RecipeDTO = &RecipeDTO{}

	err := r.db.FindOne(ctx, GetRecipeById(*oid)).Decode(recipe)

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, recipe
}

func (r *recipeRepository) CreateRecipe(recipe *RecipeNameDTO) (int, *RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := r.db.InsertOne(ctx, recipe)
	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	oid := result.InsertedID

	if oid == nil {
		log.Println("No pudo insertarse en la base de datos")
		return http.StatusInternalServerError, nil
	}

	var recipeCreated *RecipeDTO = &RecipeDTO{
		ID:   oid.(primitive.ObjectID),
		Name: recipe.Name,
	}

	return http.StatusCreated, recipeCreated
}

func (r *recipeRepository) UpdateRecipeName(oid *primitive.ObjectID, recipeName *RecipeNameDTO) (int, *RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*oid), UpdateRecipeName(*recipeName))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	var recipe *RecipeDTO = &RecipeDTO{}

	err = r.db.FindOne(ctx, GetRecipeById(*oid)).Decode(recipe)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipe
}

func (r *recipeRepository) AddIngredientToRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) (int, *RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*oid), AddIngredientToRecipe(*recipe))
	_, err = r.db.UpdateOne(ctx, GetRecipeById(*oid), SetRecipePrice())

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	var recipeUpdated *RecipeDTO = &RecipeDTO{}

	err = r.db.FindOne(ctx, GetRecipeById(*oid)).Decode(recipeUpdated)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipeUpdated
}

func (r *recipeRepository) RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) (int, *RecipeDTO) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*oid), RemoveIngredientFromRecipe(*recipe))

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	var recipeUpdated *RecipeDTO = &RecipeDTO{}

	err = r.db.FindOne(ctx, GetRecipeById(*oid)).Decode(recipeUpdated)

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, recipeUpdated
}

func (r *recipeRepository) DeleteRecipe(oid *primitive.ObjectID) (int, *primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.DeleteOne(ctx, GetRecipeById(*oid))

	if err != nil {
		log.Println(err.Error())
		return http.StatusNotFound, nil
	}

	return http.StatusOK, oid
}

func (r *recipeRepository) RemoveIngredientByPackageId(packageId *primitive.ObjectID) (int, *primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateMany(ctx, GetRecipeByPackageId(*packageId), RemovePackageFromRecipes(*packageId))

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	return http.StatusOK, packageId
}

func (r *recipeRepository) UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) (int, *primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*recipeId), SetRecipePrice())

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, recipeId
}

func (r *recipeRepository) UpdateRecipesPrice() int {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateMany(ctx, All(), SetRecipePrice())

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest
	}

	return http.StatusOK
}

func (r *recipeRepository) UpdateIngredientsPrice(packageId *primitive.ObjectID, price float64) (int, *primitive.ObjectID) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateMany(ctx, GetRecipeByPackageId(*packageId), SetIngredientPackagePrice(price), GetArrayFiltersForIngredientsByPackageId(*packageId))

	cursor, err := r.db.Find(ctx, GetRecipeByPackageId(*packageId))

	if err != nil {
		log.Println(err.Error())
		return http.StatusInternalServerError, nil
	}

	var recipes []RecipeDTO = []RecipeDTO{}

	err = cursor.All(ctx, &recipes)

	for i := 0; i < len(recipes); i++ {
		var recipePrice float64 = 0
		for j := 0; j < len(recipes[i].Ingredients); j++ {
			recipes[i].Ingredients[j].Price = recipes[i].Ingredients[j].Quantity / recipes[i].Ingredients[j].Package.Quantity * recipes[i].Ingredients[j].Package.Price
			recipePrice += recipes[i].Ingredients[j].Price
		}
		recipes[i].Price = recipePrice
		_, err = r.db.UpdateOne(ctx, GetRecipeByPackageId(*packageId), SetIngredientPrice(recipes[i]))
	}

	if err != nil {
		log.Println(err.Error())
		return http.StatusBadRequest, nil
	}

	return http.StatusOK, packageId
}
