package courses

import (
	dto "courses-api/src/dto/courses"
	"courses-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoursesController struct {
	services *services.Services
}

type CoursesControllerInterface interface {
	CreateCourse(ctx *gin.Context)
	GetAllCourses(ctx *gin.Context)
	GetCourseById(ctx *gin.Context)
	UpdateCourse(ctx *gin.Context)
	DeleteCourse(ctx *gin.Context)
}

func NewCoursesController(services *services.Services) CoursesControllerInterface {
	return &CoursesController{
		services: services,
	}
}

func (c *CoursesController) CreateCourse(ctx *gin.Context) {
	var courseDto dto.CreateCoursesRequestDto
	if err := ctx.ShouldBindJSON(&courseDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := c.services.Courses.Create(ctx, &courseDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, result)
}

func (c *CoursesController) GetAllCourses(ctx *gin.Context) {
	courses, err := c.services.Courses.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, courses)
}

func (c *CoursesController) GetCourseById(ctx *gin.Context) {
	id := ctx.Param("id")
	course, err := c.services.Courses.GetById(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, course)
}

func (c *CoursesController) UpdateCourse(ctx *gin.Context) {
	var courseDto dto.UpdateCourseDto
	if err := ctx.ShouldBindJSON(&courseDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courseDto.Id = ctx.Param("id")
	result, err := c.services.Courses.Update(ctx, &courseDto)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, result)
}

func (c *CoursesController) DeleteCourse(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := c.services.Courses.Delete(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": result})
}
