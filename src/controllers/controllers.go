package controllers

import (
	"courses-api/src/controllers/categories"
	"courses-api/src/controllers/comments"
	"courses-api/src/controllers/courses"
	"courses-api/src/controllers/ratings"
	"courses-api/src/services"
)

type Controllers struct {
	Categories categories.CategoriesControllerInterface
	Courses    courses.CoursesControllerInterface
	Comments   comments.CommentsControllerInterface
	Ratings    ratings.RatingsControllerInterface
}

func NewControllers(services *services.Services) *Controllers {
	return &Controllers{
		Categories: categories.NewCategoriesController(services),
		Courses:    courses.NewCoursesController(services),
		Comments:   comments.NewCommentsController(services),
		Ratings:    ratings.NewRatingsController(services),
	}
}
