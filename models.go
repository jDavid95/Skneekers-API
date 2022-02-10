package models

import "go.mongodb.org/mongo-driver/primitive"

type Sneaker struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Brand    string             `json:"brand,omitempty" bson:"brand,omitempty"`
	Model    string             `json:"model" bson:"model,omitempty"`
	Color    string             `json:"color" bson:"color,omitempty"`
	Year     int                `json:"year" bson:"year,omitempty"`
	Price    float32            `json:"price" bson:"price,omitempty"`
}
