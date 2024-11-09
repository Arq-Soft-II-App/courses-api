package categories

import (
	dto "courses-api/src/dto/categories"
	"courses-api/src/services/categories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoriesController struct {
	service categories.CategoryInterface
}

type CategoriesControllerInterface interface {
	CreateCategory(ctx *gin.Context)
	GetCategories(ctx *gin.Context)
}

func NewCategoriesController(service categories.CategoryInterface) CategoriesControllerInterface {
	return &CategoriesController{
		service: service,
	}
}

func (c *CategoriesController) CreateCategory(ctx *gin.Context) {
	var dto dto.CategoryDto
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if dto.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El nombre de la categoría es requerido"})
		return
	}

	if len(dto.Name) <= 4 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "El nombre de la categoría debe tener al menos 4 caracteres"})
		return
	}

	if err := c.service.Create(ctx, &dto); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

func (c *CategoriesController) GetCategories(ctx *gin.Context) {
	categories, err := c.service.GetAll(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, categories)
}
