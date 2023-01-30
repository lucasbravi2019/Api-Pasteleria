package core

import (
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertHexToObjectId(id string) (primitive.ObjectID, error) {
	log.Println("ID: " + id)
	oid, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		log.Println("Object ID invalid")
		return primitive.NewObjectID(), errors.New("object id invalid")
	}
	return oid, nil
}
