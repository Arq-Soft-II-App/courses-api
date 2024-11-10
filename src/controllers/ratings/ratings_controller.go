package ratings

import (
	dto "courses-api/src/dto/ratings"
	"courses-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RatingsController struct {
	services *services.Services
}

func NewRatingsController(services *services.Services) RatingsControllerInterface {
	return &RatingsController{
		services: services,
	}
}

type RatingsControllerInterface interface {
	NewRating(c *gin.Context)
	UpdateRating(c *gin.Context)
	GetCourseRating(c *gin.Context)
}

func (rc *RatingsController) NewRating(c *gin.Context) {
	var ratingDTO dto.RatingRequestResponseDto
	if err := c.ShouldBindJSON(&ratingDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newRating, err := rc.services.Ratings.NewRating(c.Request.Context(), &ratingDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newRating)
}

func (rc *RatingsController) UpdateRating(c *gin.Context) {
	var ratingDTO dto.RatingRequestResponseDto
	if err := c.ShouldBindJSON(&ratingDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRating, err := rc.services.Ratings.UpdateRating(c.Request.Context(), &ratingDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRating)
}

func (rc *RatingsController) GetCourseRating(c *gin.Context) {
	courseID := c.Param("id")
	rating, err := rc.services.Ratings.GetCourseRating(c.Request.Context(), courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rating)
}
