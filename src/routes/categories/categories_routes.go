package categories

import (
	"courses-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(g *gin.RouterGroup, controller controllers.CategoriesControllerInterface) {
	g.POST("/", controller.CreateCategory)
	g.GET("/", controller.GetCategories)
}
