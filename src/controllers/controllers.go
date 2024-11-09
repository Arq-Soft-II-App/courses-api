package controllers

import (
	"courses-api/src/controllers/categories"
	"courses-api/src/controllers/comments"
	"courses-api/src/controllers/courses"
	"courses-api/src/services"
)

type Controllers struct {
	Categories categories.CategoriesControllerInterface
	Courses    courses.CoursesControllerInterface
	Comments   comments.CommentsControllerInterface
}

func NewControllers(services *services.Services) *Controllers {
	return &Controllers{
		Categories: categories.NewCategoriesController(services.Categories),
		Courses:    courses.NewCoursesController(services.Courses),
		Comments:   comments.NewCommentsController(services.Comments),
	}
}
