package packages

import "go.mongodb.org/mongo-driver/bson/primitive"

type PackageDTO struct {
	ID primitive.ObjectID `bson:"_id" json:"id" validate:"required"`
}
