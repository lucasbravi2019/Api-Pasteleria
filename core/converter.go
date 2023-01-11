package core

import (
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertHexToObjectId(id string) primitive.ObjectID {
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Fatal("Object ID invalid")
	}
	return oid
}
