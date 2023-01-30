package ingredients

import "go.mongodb.org/mongo-driver/bson/primitive"

type Ingredient struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name,omitempty" validate:"required"`
	Metric   string             `bson:"metric" json:"metric,omitempty" validate:"required"`
	Quantity int                `bson:"quantity" json:"quantity,omitempty" validate:"required"`
	Price    float32            `bson:"price" json:"price,omitempty" validate:"required"`
}
