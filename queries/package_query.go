package queries

import (
	"github.com/lucasbravi2019/pasteleria/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPackageById(packageId primitive.ObjectID) bson.M {
	return bson.M{"_id": packageId}
}

func UpdatePackageById(body models.Package) bson.M {
	return bson.M{"$set": bson.M{
		"metric":   body.Metric,
		"quantity": body.Quantity,
	}}
}
