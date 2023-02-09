package ingredients

import (
	"strings"

	"github.com/lucasbravi2019/pasteleria/api/packages"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetIngredientById(id primitive.ObjectID) bson.M {
	return bson.M{
		"_id": id,
	}
}

func GetIngredientByPackageId(packageId primitive.ObjectID) bson.M {
	return bson.M{
		"packages._id": packageId,
	}
}

func GetAggregateShowIngredients() []bson.M {
	packagesUnwind := bson.M{"$unwind": "$packages"}
	packagesLookup := bson.M{"$lookup": bson.M{
		"from":         "packages",
		"localField":   "packages._id",
		"foreignField": "_id",
		"as":           "package",
	}}
	packageUnwind := bson.M{"$unwind": "$package"}
	priceSet := bson.M{"$set": bson.M{
		"package.price": "$packages.price",
	}}
	packagesUnset := bson.M{"$unset": "packages"}
	group := bson.M{"$group": bson.M{
		"_id": "$_id",
		"name": bson.M{
			"$first": "$name",
		},
		"package": bson.M{
			"$push": "$package",
		},
	}}

	return []bson.M{packagesUnwind, packagesLookup, packageUnwind, priceSet, packagesUnset, group}
}

func GetAggregateCreateIngredients(ingredient *Ingredient) []bson.D {
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

	return []bson.D{project, match}
}

func SetIngredientPackages(envase packages.Package) bson.M {
	return bson.M{"$set": bson.M{
		"packages.$": envase,
	}}
}

func GetIngredientWithoutExistingPackage(ingredientOid primitive.ObjectID, packageOid primitive.ObjectID) bson.D {
	return bson.D{{"_id", ingredientOid}, {"packages._id", bson.D{{"$ne", packageOid}}}}
}

func PushPackageIntoIngredient(envase packages.Package) bson.M {
	return bson.M{"$push": bson.M{
		"packages": envase,
	}}
}
