package clients

import (
	"courses-api/src/clients/categories"
	"courses-api/src/clients/courses"

	"go.mongodb.org/mongo-driver/mongo"
)

type Clients struct {
	Categories *categories.CategoriesClient
	Courses    *courses.CourseClient
}

func NewClients(db *mongo.Database) *Clients {
	return &Clients{
		Categories: categories.NewCategoriesClient(db),
		Courses:    courses.NewCourseClient(db),
	}
}
