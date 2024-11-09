package courses

import (
	"courses-api/src/controllers/courses"

	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.RouterGroup, controller courses.CoursesControllerInterface) {
	g.POST("/", controller.CreateCourse)
	g.GET("/", controller.GetAllCourses)
	g.GET("/:id", controller.GetCourseById)
	g.PUT("/:id", controller.UpdateCourse)
	g.DELETE("/:id", controller.DeleteCourse)
}
