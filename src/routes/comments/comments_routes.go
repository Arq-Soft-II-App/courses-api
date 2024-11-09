package comments

import (
	"courses-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func CommentsRoutes(g *gin.RouterGroup, controller controllers.Controllers) {
	g.POST("/", controller.Comments.NewComment)
	g.GET("/:course_id", controller.Comments.GetCourseComments)
	g.PUT("/", controller.Comments.UpdateComment)
}
