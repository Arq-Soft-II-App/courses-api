package categories

import (
	"courses-api/src/controllers/categories"

	"github.com/gin-gonic/gin"
)

func CategoriesRoutes(g *gin.RouterGroup, controller categories.CategoriesControllerInterface) {
	g.POST("/", controller.CreateCategory)
	g.GET("/", controller.GetCategories)
}
