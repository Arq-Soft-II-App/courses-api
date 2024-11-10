package comments

import (
	dto "courses-api/src/dto/comments"
	"courses-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsController struct {
	services *services.Services
}

type CommentsControllerInterface interface {
	NewComment(c *gin.Context)
	GetCourseComments(c *gin.Context)
	UpdateComment(c *gin.Context)
}

func NewCommentsController(services *services.Services) CommentsControllerInterface {
	return &CommentsController{
		services: services,
	}
}

func (cc *CommentsController) NewComment(c *gin.Context) {
	var commentDTO dto.CommentsDto
	if err := c.ShouldBindJSON(&commentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment, err := cc.services.Comments.NewComment(c.Request.Context(), &commentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newComment)
}

func (cc *CommentsController) GetCourseComments(c *gin.Context) {
	courseID := c.Param("course_id")
	comments, err := cc.services.Comments.GetCourseComments(c.Request.Context(), courseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (cc *CommentsController) UpdateComment(c *gin.Context) {
	var commentDTO dto.CommentsDto
	if err := c.ShouldBindJSON(&commentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedComment, err := cc.services.Comments.UpdateComment(c.Request.Context(), &commentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedComment)
}
