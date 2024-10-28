package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Rating struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	CourseID  primitive.ObjectID `bson:"course_id"`
	UserID    string             `bson:"user_id"`
	Rating    int                `bson:"rating"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type Ratings []Rating
