package recipes

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetRecipeById(oid primitive.ObjectID) bson.M {
	return bson.M{"_id": oid}
}

func UpdateRecipeName(dto RecipeNameDTO) bson.M {
	return bson.M{"$set": bson.M{"name": dto.Name}}
}

func AddIngredientToRecipe(recipe RecipeIngredient) bson.M {
	return bson.M{"$addToSet": bson.M{"$ingredients": recipe}}
}

func GetAggregateAllRecipe() mongo.Pipeline {
	unwindIngredients := bson.D{{"$unwind", "$ingredients"}}

	lookupIngredients := bson.D{{"$lookup", bson.D{{"from", "ingredients"}, {"localField", "ingredients._id"}, {"foreignField", "_id"}, {"as", "ingredient"}}}}

	unwindIngredient := bson.D{{"$unwind", "$ingredient"}}

	lookupPackages := bson.D{{"$lookup", bson.D{{"from", "packages"}, {"localField", "ingredients.package._id"}, {"foreignField", "_id"}, {"as", "ingredient.packages"}}}}

	unwindPackages := bson.D{{"$unwind", "$ingredient.packages"}}

	setFields := bson.D{{"$set", bson.D{{"ingredient.price", "$ingredients.price"}, {"ingredient.quantity", "$ingredients.quantity"}}}}

	group := bson.D{{"$group", bson.D{
		{"_id", "$_id"},
		{"name", bson.D{{"$first", "$name"}}}, {"ingredients", bson.D{{"$push", "$ingredient"}}},
		{"price", bson.D{{"$sum", "$ingredient.price"}}}}}}

	return mongo.Pipeline{unwindIngredients, lookupIngredients, unwindIngredient, lookupPackages, unwindPackages, setFields, group}
}

func GetAggregateRecipeById(oid primitive.ObjectID) mongo.Pipeline {
	matchById := bson.D{{"$match", bson.D{{"_id", oid}}}}

	unwindIngredients := bson.D{{"$unwind", "$ingredients"}}

	lookupIngredients := bson.D{{"$lookup", bson.D{{"from", "ingredients"}, {"localField", "ingredients._id"}, {"foreignField", "_id"}, {"as", "ingredient"}}}}

	unwindIngredient := bson.D{{"$unwind", "$ingredient"}}

	lookupPackages := bson.D{{"$lookup", bson.D{{"from", "packages"}, {"localField", "ingredients.package._id"}, {"foreignField", "_id"}, {"as", "ingredient.packages"}}}}

	unwindPackages := bson.D{{"$unwind", "$ingredient.packages"}}

	setFields := bson.D{{"$set", bson.D{{"ingredient.price", "$ingredients.price"}, {"ingredient.quantity", "$ingredients.quantity"}}}}

	group := bson.D{{"$group", bson.D{
		{"_id", "$_id"},
		{"name", bson.D{{"$first", "$name"}}}, {"ingredients", bson.D{{"$push", "$ingredient"}}},
		{"price", bson.D{{"$sum", "$ingredient.price"}}}}}}

	return mongo.Pipeline{matchById, unwindIngredients, lookupIngredients, unwindIngredient, lookupPackages, unwindPackages, setFields, group}
}
