package controllers

import (
	"courses-api/src/controllers/categories"
	"courses-api/src/controllers/courses"
	"courses-api/src/services"
)

type CategoriesControllerInterface = categories.CategoriesControllerInterface
type CoursesControllerInterface = courses.CoursesControllerInterface

type Controllers struct {
	Categories CategoriesControllerInterface
	Courses    CoursesControllerInterface
}

func NewControllers(services *services.Services) *Controllers {
	return &Controllers{
		Categories: categories.NewCategoriesController(services.Categories),
		Courses:    courses.NewCoursesController(services.Courses),
	}
}
