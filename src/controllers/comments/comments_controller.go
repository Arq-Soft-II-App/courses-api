package comments

import (
	dto "courses-api/src/dto/comments"
	"courses-api/src/services/comments"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsController struct {
	service comments.CommentsInterface
}

func NewCommentsController(service comments.CommentsInterface) *CommentsController {
	return &CommentsController{
		service: service,
	}
}

type CommentsControllerInterface interface {
	NewComment(c *gin.Context)
	GetCourseComments(c *gin.Context)
	UpdateComment(c *gin.Context)
}

func (cc *CommentsController) NewComment(c *gin.Context) {
	var commentDTO dto.CommentsDto
	if err := c.ShouldBindJSON(&commentDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment, err := cc.service.NewComment(c.Request.Context(), &commentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newComment)
}

func (cc *CommentsController) GetCourseComments(c *gin.Context) {
	courseID := c.Param("course_id")
	comments, err := cc.service.GetCourseComments(c.Request.Context(), courseID)
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

	updatedComment, err := cc.service.UpdateComment(c.Request.Context(), &commentDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedComment)
}
