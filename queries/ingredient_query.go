package queries

import (
	"strings"

	"github.com/lucasbravi2019/pasteleria/dto"
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetIngredientById(oid primitive.ObjectID) bson.M {
	return bson.M{"_id": oid}
}

func GetIngredientByPackageId(packageId primitive.ObjectID) bson.M {
	return bson.M{
		"packages._id": packageId,
	}
}

func GetAggregateCreateIngredients(ingredient *dto.IngredientNameDTO) mongo.Pipeline {
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

func UpdateIngredientName(dto dto.IngredientNameDTO) bson.M {
	return bson.M{"$set": bson.M{"name": dto.Name}}
}

func PushPackageIntoIngredient(envase models.IngredientPackage) bson.M {
	return bson.M{"$addToSet": bson.M{
		"packages": envase,
	}}
}

func PullPackageFromIngredients(envase dto.IngredientPackageDTO) bson.M {
	return bson.M{"$pull": bson.M{"packages": bson.M{"_id": envase.PackageOid}}}
}

func SetIngredientPrice(price float64) bson.M {
	return bson.M{
		"$set": bson.M{
			"packages.$[package].price": price,
		},
	}
}

func GetArrayFilterForPackageId(oid primitive.ObjectID) *options.UpdateOptions {
	return options.Update().SetArrayFilters(options.ArrayFilters{
		Filters: []interface{}{
			bson.M{
				"package._id": oid,
			},
		},
	})
}
