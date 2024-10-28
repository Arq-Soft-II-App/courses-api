package routes

import (
	"courses-api/src/controllers"
	"courses-api/src/routes/categories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, controller *controllers.Controllers) {

	categoriesRoutes := router.Group("/categories")
	{
		categories.CategoriesRoutes(categoriesRoutes, controller.Categories)
	}

}
