package recipes

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type recipeRepository struct {
	db *mongo.Collection
}

type RecipeRepository interface {
	FindAllRecipes() *[]RecipeDTO
	FindRecipeByOID(oid *primitive.ObjectID) *RecipeDTO
	FindRecipesByPackageId(oid *primitive.ObjectID) []RecipeDTO
	CreateRecipe(recipe *RecipeNameDTO) *primitive.ObjectID
	UpdateRecipeName(oid *primitive.ObjectID, recipeName *RecipeNameDTO) error
	AddIngredientToRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) error
	RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) error
	DeleteRecipe(oid *primitive.ObjectID) error
	RemoveIngredientByPackageId(packageId *primitive.ObjectID) error
	UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error
	UpdateIngredientPackagePrice(packageId *primitive.ObjectID, price float64) error
	UpdateIngredientsPrice(packageId *primitive.ObjectID, recipe RecipeDTO) error
	UpdateRecipesPrice() error
}

var recipeRepositoryInstance *recipeRepository

func (r *recipeRepository) FindAllRecipes() *[]RecipeDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := r.db.Find(ctx, All())

	var recipes *[]RecipeDTO = &[]RecipeDTO{}

	if err != nil {
		log.Println(err.Error())
		return recipes
	}

	err = cursor.All(ctx, recipes)

	if err != nil {
		log.Println(err.Error())
	}

	return recipes
}

func (r *recipeRepository) FindRecipeByOID(oid *primitive.ObjectID) *RecipeDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var recipe *RecipeDTO = &RecipeDTO{}

	err := r.db.FindOne(ctx, GetRecipeById(*oid)).Decode(recipe)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	return recipe
}

func (r *recipeRepository) FindRecipesByPackageId(packageId *primitive.ObjectID) []RecipeDTO {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	cursor, err := r.db.Find(ctx, GetRecipeByPackageId(*packageId))

	var recipes []RecipeDTO = []RecipeDTO{}

	if err != nil {
		log.Println(err.Error())
		return recipes
	}

	err = cursor.All(ctx, &recipes)

	if err != nil {
		log.Println(err.Error())
	}

	return recipes
}

func (r *recipeRepository) CreateRecipe(recipe *RecipeNameDTO) *primitive.ObjectID {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := r.db.InsertOne(ctx, recipe)

	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if result.InsertedID == nil {
		return nil
	}

	id := result.InsertedID.(primitive.ObjectID)

	return &id
}

func (r *recipeRepository) UpdateRecipeName(oid *primitive.ObjectID, recipeName *RecipeNameDTO) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*oid), UpdateRecipeName(*recipeName))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *recipeRepository) AddIngredientToRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*oid), AddIngredientToRecipe(*recipe))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *recipeRepository) RemoveIngredientFromRecipe(oid *primitive.ObjectID, recipe *RecipeIngredient) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*oid), RemoveIngredientFromRecipe(*recipe))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *recipeRepository) DeleteRecipe(oid *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.DeleteOne(ctx, GetRecipeById(*oid))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *recipeRepository) RemoveIngredientByPackageId(packageId *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.TODO(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateMany(ctx, GetRecipeByPackageId(*packageId), RemovePackageFromRecipes(*packageId))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *recipeRepository) UpdateRecipeByIdPrice(recipeId *primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeById(*recipeId), SetRecipePrice())

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *recipeRepository) UpdateRecipesPrice() error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateMany(ctx, All(), SetRecipePrice())

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *recipeRepository) UpdateIngredientPackagePrice(packageId *primitive.ObjectID, price float64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateMany(ctx, GetRecipeByPackageId(*packageId), SetIngredientPackagePrice(price), GetArrayFiltersForIngredientsByPackageId(*packageId))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}

func (r *recipeRepository) UpdateIngredientsPrice(packageId *primitive.ObjectID, recipe RecipeDTO) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	_, err := r.db.UpdateOne(ctx, GetRecipeByPackageId(*packageId), SetIngredientPrice(recipe))

	if err != nil {
		log.Println(err.Error())
	}

	return err
}
