package courses

import (
	dto "courses-api/src/dto/courses"
	"courses-api/src/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CoursesController struct {
	services *services.Services
}

type CoursesControllerInterface interface {
	CreateCourse(ctx *gin.Context)
	GetAllCourses(ctx *gin.Context)
	GetCourseList(ctx *gin.Context)
	GetCourseById(ctx *gin.Context)
	UpdateCourse(ctx *gin.Context)
	DeleteCourse(ctx *gin.Context)
}

func NewCoursesController(services *services.Services) CoursesControllerInterface {
	return &CoursesController{
		services: services,
	}
}

func (c *CoursesController) GetCourseList(ctx *gin.Context) {
	var requestBody struct {
		IDs []string `json:"ids" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Los IDs son requeridos y deben ser un array de strings"})
		return
	}

	if len(requestBody.IDs) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "La lista de IDs no puede estar vacía"})
		return
	}
	fmt.Println("IDs:", requestBody.IDs)
	courses, err := c.services.Courses.GetCourseList(ctx.Request.Context(), requestBody.IDs)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener cursos"})
		return
	}

	ctx.JSON(http.StatusOK, courses)
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
