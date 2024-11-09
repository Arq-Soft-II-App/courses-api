package clients

import (
	"courses-api/src/clients/categories"
	"courses-api/src/clients/comments"
	"courses-api/src/clients/courses"

	"go.mongodb.org/mongo-driver/mongo"
)

type Clients struct {
	Categories categories.CategoryClientInterface
	Courses    courses.CourseClientInterface
	Comments   comments.CommentsClientInterface
}

func NewClients(db *mongo.Database) *Clients {
	return &Clients{
		Categories: categories.NewCategoriesClient(db),
		Courses:    courses.NewCourseClient(db),
		Comments:   comments.NewCommentsClient(db),
	}
}
