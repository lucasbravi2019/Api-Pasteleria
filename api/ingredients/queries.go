package ingredients

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetIngredientByIdAggregation(id primitive.ObjectID) mongo.Pipeline {
	match := bson.D{{"$match", bson.D{{"_id", id}}}}
	packagesUnwind := bson.D{{"$unwind", bson.D{{"path", "$packages"}, {"preserveNullAndEmptyArrays", true}}}}
	packagesLookup := bson.D{{"$lookup", bson.D{
		{"from", "packages"},
		{"localField", "packages._id"},
		{"foreignField", "_id"},
		{"as", "package"},
	}}}
	packageUnwind := bson.D{{"$unwind", bson.D{{"path", "$package"}, {"preserveNullAndEmptyArrays", true}}}}

	priceSet := bson.D{{"$set", bson.D{{"package.price", "$packages.price"}}}}

	packagesUnset := bson.D{{"$unset", "packages"}}

	group := bson.D{{"$group", bson.D{{"_id", "$_id"}, {"name", bson.D{{"$first", "$name"}}}, {"package", bson.D{{"$push", "$package"}}}}}}

	addFields := bson.D{{"$addFields", bson.D{{"package", bson.D{{"$filter", bson.D{{"input", "$package"},
		{"cond", bson.D{{"$ne", bson.A{"$$this._id", "undefined"}}}}}}}}}}}

	return mongo.Pipeline{match, packagesUnwind, packagesLookup, packageUnwind, priceSet, packagesUnset, group, addFields}
}

func GetIngredientById(oid primitive.ObjectID) bson.M {
	return bson.M{"_id": oid}
}

func GetIngredientByPackageId(packageId primitive.ObjectID) bson.M {
	return bson.M{
		"packages._id": packageId,
	}
}

func GetAllIngredients() mongo.Pipeline {
	packagesUnwind := bson.D{{"$unwind", bson.D{{"path", "$packages"}, {"preserveNullAndEmptyArrays", true}}}}
	packagesLookup := bson.D{{"$lookup", bson.D{
		{"from", "packages"},
		{"localField", "packages._id"},
		{"foreignField", "_id"},
		{"as", "package"},
	}}}
	packageUnwind := bson.D{{"$unwind", bson.D{{"path", "$package"}, {"preserveNullAndEmptyArrays", true}}}}

	priceSet := bson.D{{"$set", bson.D{{"package.price", "$packages.price"}}}}

	group := bson.D{{"$group", bson.D{{"_id", "$_id"}, {"name", bson.D{{"$first", "$name"}}},
		{"packages", bson.D{{"$push", bson.D{{"$cond", bson.A{bson.D{{"$ne", bson.A{"$package._id", "$packages._id"}}}, "$$REMOVE", "$package"}}}}}}}}}

	return mongo.Pipeline{packagesUnwind, packagesLookup, packageUnwind, priceSet, group}
}

func GetAggregateCreateIngredients(ingredient *IngredientNameDTO) mongo.Pipeline {
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

func GetIngredientWithoutExistingPackage(ingredientOid primitive.ObjectID, packageOid primitive.ObjectID) bson.D {
	return bson.D{{"_id", ingredientOid}, {"packages._id", bson.D{{"$ne", packageOid}}}}
}

func UpdateIngredientName(dto IngredientNameDTO) bson.M {
	return bson.M{"$set": bson.M{"name": dto.Name}}
}

func PushPackageIntoIngredient(envase IngredientPackage) bson.M {
	return bson.M{"$addToSet": bson.M{
		"packages": envase,
	}}
}

func PullPackageFromIngredients(envase IngredientPackageDTO) bson.M {
	return bson.M{"$pull": bson.M{"packages": bson.M{"_id": envase.PackageOid}}}
}

func SetIngredientPrice(price float64) bson.M {
	return bson.M{
		"$set": bson.M{
			"packages.$[package].package.price": price,
		},
	}
}

func GetArrayFilterForPackageId(oid primitive.ObjectID) *options.UpdateOptions {
	return options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{
				"package.package._id": oid,
			},
		},
	})
}
