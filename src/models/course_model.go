package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Course struct {
	Id                primitive.ObjectID `bson:"_id,omitempty"`
	CourseName        string             `bson:"course_name"`
	CourseDescription string             `bson:"description"`
	CoursePrice       float64            `bson:"price"`
	CourseDuration    int                `bson:"duration"`
	CourseInitDate    string             `bson:"init_date"`
	CourseState       bool               `bson:"state" default:"true"`
	CourseCapacity    int                `bson:"cupo" default:"15"`
	CourseImage       string             `bson:"image" default:"https://upload.wikimedia.org/wikipedia/commons/a/a3/Image-not-found.png"`
	CategoryID        primitive.ObjectID `bson:"category_id" ref:"categories"`
	CategoryName      string             `bson:"category_name,omitempty" json:"category_name,omitempty"`
	RatingAvg         float64            `bson:"ratingavg,omitempty" json:"ratingavg,omitempty"`
}

type Courses []Course
