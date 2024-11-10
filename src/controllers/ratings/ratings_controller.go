package ratings

import (
	dto "courses-api/src/dto/ratings"
	"courses-api/src/services/ratings"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RatingsController struct {
	service ratings.RatingsInterface
}

func NewRatingsController(service ratings.RatingsInterface) *RatingsController {
	return &RatingsController{
		service: service,
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

	newRating, err := rc.service.NewRating(c.Request.Context(), &ratingDTO)
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

	updatedRating, err := rc.service.UpdateRating(c.Request.Context(), &ratingDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedRating)
}

func (rc *RatingsController) GetCourseRating(c *gin.Context) {
	courseID := c.Param("id")
	rating, err := rc.service.GetCourseRating(c.Request.Context(), courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rating)
}
