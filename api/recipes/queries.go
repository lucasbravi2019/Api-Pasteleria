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
	return bson.M{"$addToSet": bson.M{"ingredients": recipe}}
}

func GetAggregateAllRecipe() mongo.Pipeline {
	unwindIngredients := bson.D{{"$unwind", bson.D{{"path", "$ingredients"}, {"preserveNullAndEmptyArrays", true}}}}

	lookupIngredients := bson.D{{"$lookup", bson.D{{"from", "ingredients"}, {"localField", "ingredients._id"}, {"foreignField", "_id"}, {"as", "ingredient"}}}}

	unwindIngredient := bson.D{{"$unwind", bson.D{{"path", "$ingredient"}, {"preserveNullAndEmptyArrays", true}}}}

	unwindIngredientPackages := bson.D{{"$unwind", "$ingredient.packages"}}

	setPrice := bson.D{{"$set", bson.D{{"ingredients.price", "$ingredient.packages.price"}}}}

	lookupPackages := bson.D{{"$lookup", bson.D{{"from", "packages"}, {"localField", "ingredients.package._id"}, {"foreignField", "_id"}, {"as", "ingredient.packages"}}}}

	unwindPackages := bson.D{{"$unwind", "$ingredient.packages"}}

	setFields := bson.D{{"$set", bson.D{
		{"ingredient.packages.price", "$ingredients.price"},
		{"ingredient.price", bson.D{
			{"$multiply", bson.A{bson.D{{"$divide", bson.A{"$ingredients.quantity", "$ingredient.packages.quantity"}}}, "$ingredients.price"}}}}}}}

	group := bson.D{{"$group", bson.D{
		{"_id", "$_id"},
		{"name", bson.D{{"$first", "$name"}}}, {"ingredients", bson.D{{"$push", "$ingredient"}}},
		{"price", bson.D{{"$sum", "$ingredient.price"}}}}}}

	return mongo.Pipeline{unwindIngredients, lookupIngredients, unwindIngredient, unwindIngredientPackages, setPrice, lookupPackages, unwindPackages, setFields, group}
}

func GetAggregateRecipeById(oid primitive.ObjectID) mongo.Pipeline {
	matchById := bson.D{{"$match", bson.D{{"_id", oid}}}}

	unwindIngredients := bson.D{{"$unwind", bson.D{{"path", "$ingredients"}, {"preserveNullAndEmptyArrays", true}}}}

	lookupIngredients := bson.D{{"$lookup", bson.D{{"from", "ingredients"}, {"localField", "ingredients._id"}, {"foreignField", "_id"}, {"as", "ingredient"}}}}

	unwindIngredient := bson.D{{"$unwind", bson.D{{"path", "$ingredient"}, {"preserveNullAndEmptyArrays", true}}}}

	lookupPackages := bson.D{{"$lookup", bson.D{{"from", "packages"}, {"localField", "ingredients.package._id"}, {"foreignField", "_id"}, {"as", "ingredient.packages"}}}}

	unwindPackages := bson.D{{"$unwind", bson.D{{"path", "$ingredient.packages"}, {"preserveNullAndEmptyArrays", true}}}}

	setFields := bson.D{{"$set", bson.D{
		{"ingredient.quantity", "$ingredients.quantity"},
		{"ingredient.packages.price", "$ingredients.package.price"},
		{"ingredient.price", bson.D{
			{"$multiply", bson.A{bson.D{{"$divide", bson.A{"$ingredients.quantity", "$ingredient.packages.quantity"}}}, "$ingredients.package.price"}}}}}}}

	group := bson.D{{"$group", bson.D{
		{"_id", "$_id"},
		{"name", bson.D{{"$first", "$name"}}}, {"ingredients", bson.D{{"$push", "$ingredient"}}},
		{"price", bson.D{{"$sum", "$ingredient.price"}}}}}}

	return mongo.Pipeline{matchById, unwindIngredients, lookupIngredients, unwindIngredient, lookupPackages, unwindPackages, setFields, group}
}
