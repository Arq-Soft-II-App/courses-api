package clients

import (
	"courses-api/src/clients/categories"
	"courses-api/src/clients/comments"
	"courses-api/src/clients/courses"
	"courses-api/src/clients/ratings"

	"go.mongodb.org/mongo-driver/mongo"
)

type Clients struct {
	Categories categories.CategoryClientInterface
	Courses    courses.CourseClientInterface
	Comments   comments.CommentsClientInterface
	Ratings    ratings.RatingsClientInterface
}

func NewClients(db *mongo.Database) *Clients {
	return &Clients{
		Categories: categories.NewCategoriesClient(db),
		Courses:    courses.NewCourseClient(db),
		Comments:   comments.NewCommentsClient(db),
		Ratings:    ratings.NewRatingsClient(db),
	}
}
