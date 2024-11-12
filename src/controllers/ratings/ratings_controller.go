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
	GetAllRatings(c *gin.Context)
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

func (rc *RatingsController) GetAllRatings(c *gin.Context) {
	rating, err := rc.services.Ratings.GetAllRatings(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rating)
}
