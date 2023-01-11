package recipes

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name        string             `bson:"recipeName" json:"recipeName,omitempty"`
	Ingredients []Ingredient       `bson:"ingredients" json:"ingredients,omitempty"`
	Price       float32            `bson:"price" json:"price,omitempty"`
}

type Ingredient struct {
	ID     primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name   string             `bson:"name" json:"name,omitempty"`
	Metric string             `bson:"metric" json:"metric,omitempty"`
	Price  float32            `bson:"price" json:"price,omitempty"`
}
