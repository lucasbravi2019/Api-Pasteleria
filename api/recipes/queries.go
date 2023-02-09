package recipes

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRecipeById(oid primitive.ObjectID) bson.M {
	return bson.M{"_id": oid}
}

func UpdateRecipe(recipe Recipe) bson.M {
	return bson.M{"$set": recipe}
}
