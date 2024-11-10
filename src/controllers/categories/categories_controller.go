package categories

import (
	dto "courses-api/src/dto/categories"
	appErrors "courses-api/src/errors"
	"courses-api/src/services"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	services *services.Services
}

type CategoriesControllerInterface interface {
	CreateCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
}

func NewCategoriesController(services *services.Services) CategoriesControllerInterface {
	return &CategoriesController{
		services: services,
	}
}

func (c *CategoriesController) CreateCategory(ctx *gin.Context) {
	var categoryDto dto.CategoryDto
	if err := ctx.ShouldBindJSON(&categoryDto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if categoryDto.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El nombre de la categoría es requerido"})
		return
	}

	if len(categoryDto.Name) <= 4 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El nombre de la categoría debe tener al menos 4 caracteres"})
		return
	}

	_, err := c.services.Categories.Create(ctx, &categoryDto)
	if err != nil {
		var appErr *appErrors.Error
		if errors.As(err, &appErr) {
			ctx.JSON(appErr.HTTPStatusCode, gin.H{"error": appErr.Message})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": `Categoría creada exitosamente`})
}

func (c *CategoriesController) GetCategories(ctx *gin.Context) {
	categories, err := c.services.Categories.GetAll(ctx)
	if err != nil {
		var appErr *appErrors.Error
		if errors.As(err, &appErr) {
			ctx.JSON(appErr.HTTPStatusCode, gin.H{"error": appErr.Message})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno del servidor"})
		}
		return
	}

	ctx.JSON(http.StatusOK, categories)
}
