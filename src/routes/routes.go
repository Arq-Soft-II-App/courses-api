package routes

import (
	"courses-api/src/controllers"
	"courses-api/src/routes/categories"
	"courses-api/src/routes/courses"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, controller *controllers.Controllers) {

	categoriesRoutes := router.Group("/categories")
	{
		categories.CategoriesRoutes(categoriesRoutes, controller.Categories)
	}

	coursesRoutes := router.Group("/courses")
	{
		courses.CoursesRoutes(coursesRoutes, controller.Courses)
	}

}
