package courses

import (
	"courses-api/src/controllers"

	"github.com/gin-gonic/gin"
)

func CoursesRoutes(g *gin.RouterGroup, controller controllers.Controllers) {
	g.POST("/", controller.Courses.CreateCourse)
	g.GET("/", controller.Courses.GetAllCourses)
	g.POST("/getCourseList", controller.Courses.GetCourseList)
	g.GET("/:id", controller.Courses.GetCourseById)
	g.PUT("/:id", controller.Courses.UpdateCourse)
	g.DELETE("/:id", controller.Courses.DeleteCourse)

}
