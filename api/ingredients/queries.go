package ingredients

import (
	"strings"

	"github.com/lucasbravi2019/pasteleria/api/packages"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func GetAggregateShowIngredients() mongo.Pipeline {
	packagesUnwind := bson.D{{"$unwind", "$packages"}}
	packagesLookup := bson.D{{"$lookup", bson.M{
		"from":         "packages",
		"localField":   "packages._id",
		"foreignField": "_id",
		"as":           "package",
	}}}
	packageUnwind := bson.D{{"$unwind", "$package"}}

	priceSet := bson.D{{"$set", bson.M{
		"package.price": "$packages.price",
	}}}
	packagesUnset := bson.D{{"$unset", "packages"}}

	group := bson.D{{"$group", bson.M{
		"_id": "$_id",
		"name": bson.M{
			"$first": "$name",
		},
		"package": bson.M{
			"$push": "$package",
		},
	}}}

	return mongo.Pipeline{packagesUnwind, packagesLookup, packageUnwind, priceSet, packagesUnset, group}
}

func GetAggregateCreateIngredients(ingredient *IngredientDTO) mongo.Pipeline {
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

	return mongo.Pipeline{project, match}
}

func SetIngredientPackages(envase packages.Package) bson.M {
	return bson.M{"$set": bson.M{
		"packages.$": envase,
	}}
}

func GetIngredientWithoutExistingPackage(ingredientOid primitive.ObjectID, packageOid primitive.ObjectID) bson.D {
	return bson.D{{"_id", ingredientOid}, {"packages._id", bson.D{{"$ne", packageOid}}}}
}

func UpdateIngredientName(dto IngredientDTO) bson.M {
	return bson.M{"$set": bson.M{"name": dto.Name}}
}

func PushPackageIntoIngredient(dto IngredientPackageDTO) bson.M {
	return bson.M{"$addToSet": bson.M{
		"packages": bson.M{
			"_id":   dto.PackageOid,
			"price": dto.Price,
		},
	}}
}
