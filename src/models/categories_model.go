package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Category_Name string             `bson:"category_name" json:"category_name" validate:"required"`
}
