package ingredients

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ingredient struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name,omitempty"`
	Metric string             `bson:"metric" json:"metric,omitempty"`
	Price  float32            `bson:"price" json:"price,omitempty"`
}
