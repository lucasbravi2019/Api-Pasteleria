package core

import (
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConvertUrlVarToObjectId(param string, c *gin.Context) (*primitive.ObjectID, error) {
	id, err := GetUrlVars(c, param)

	if err != nil {
		return nil, err
	}

	return ConvertToObjectId(id)
}

func ConvertUrlParamToObjectId(param string, c *gin.Context) (*primitive.ObjectID, error) {
	id, err := GetUrlParams(c, param)

	if err != nil {
		return nil, err
	}

	return ConvertToObjectId(id)
}

func ConvertToObjectId(param string) (*primitive.ObjectID, error) {
	oid, err := primitive.ObjectIDFromHex(param)

	if err != nil {
		log.Println("Object ID invalid")
		return nil, err
	}

	return &oid, nil
}
