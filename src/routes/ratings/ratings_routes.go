package ratings

import (
	"courses-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func RatingsRoutes(g *gin.RouterGroup, controller controllers.Controllers) {
	g.POST("/", controller.Ratings.NewRating)
	g.GET("/", controller.Ratings.GetAllRatings)
	g.PUT("/", controller.Ratings.UpdateRating)
}
