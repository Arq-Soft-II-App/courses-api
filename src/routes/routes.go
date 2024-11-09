package routes

import (
	"courses-api/src/controllers"
	"courses-api/src/routes/categories"
	"courses-api/src/routes/comments"
	"courses-api/src/routes/courses"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, controller controllers.Controllers) {
	categoriesRoutes := router.Group("/courses/categories")
	{
		categories.CategoriesRoutes(categoriesRoutes, controller)
	}

	commentsRoutes := router.Group("/courses/comments")
	{
		comments.CommentsRoutes(commentsRoutes, controller)
	}

	coursesRoutes := router.Group("/courses")
	{
		courses.CoursesRoutes(coursesRoutes, controller)
	}

}
