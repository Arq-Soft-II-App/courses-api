package services

import (
	"courses-api/src/clients"
	"courses-api/src/services/categories"
	"courses-api/src/services/comments"
	"courses-api/src/services/courses"
)

type Services struct {
	Categories categories.CategoryInterface
	Courses    courses.CourseInterface
	Comments   comments.CommentsInterface
}

func NewServices(clients *clients.Clients) *Services {
	return &Services{
		Categories: categories.NewCategoriesService(clients.Categories),
		Courses:    courses.NewCoursesService(clients.Courses),
		Comments:   comments.NewCommentsService(clients.Comments),
	}
}
