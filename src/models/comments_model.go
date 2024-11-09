package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Comment struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Text     string             `bson:"text"`
	UserId   string             `bson:"user_id"`
	CourseId primitive.ObjectID `bson:"course_id" ref:"courses"`
}

type Comments []Comment
