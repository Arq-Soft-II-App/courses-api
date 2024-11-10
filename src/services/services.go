package services

import (
	"courses-api/src/clients"
	"courses-api/src/services/categories"
	"courses-api/src/services/comments"
	"courses-api/src/services/courses"
	"courses-api/src/services/ratings"
)

type Services struct {
	Categories categories.CategoryInterface
	Courses    courses.CourseInterface
	Comments   comments.CommentsInterface
	Ratings    ratings.RatingsInterface
}

func NewServices(clients *clients.Clients) *Services {
	return &Services{
		Categories: categories.NewCategoriesService(clients),
		Courses:    courses.NewCoursesService(clients),
		Comments:   comments.NewCommentsService(clients),
		Ratings:    ratings.NewRatingsService(clients),
	}
}
