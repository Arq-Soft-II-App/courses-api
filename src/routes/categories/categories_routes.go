package categories

import (
	"courses-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(g *gin.RouterGroup, controller controllers.Controllers) {
	g.POST("/", controller.Categories.CreateCategory)
	g.GET("/", controller.Categories.GetCategories)
}
