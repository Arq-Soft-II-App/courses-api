package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	CourseID primitive.ObjectID `bson:"course_id"`
	UserID   string             `bson:"user_id"`
	Rating   int                `bson:"rating"`
}

type Ratings []Rating
