package comments

/* import (
	dto "courses-api/src/dto/comments"
	"courses-api/src/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CommentsController struct {
	service *services.CommentsServiceInterface
}

func NewCommentsController(service *services.CommentsServiceInterface) *CommentsController {
	return &CommentsController{
		service: service,
	}
}

type CommentsControllerInterface interface {
	NewComment(c *gin.Context)
	GetCourseComments(c *gin.Context)
}

func (c *CommentsController) NewComment(ctx *gin.Context) {
	newCommentDTO := &dto.CommentsDto{}
	if err := ctx.BindJSON(newCommentDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
			"code":  "INVALID_REQUEST",
		})
	}

	comment, err := c.service.NewComment(*newCommentDTO)
	if err != nil {
		if customError, ok := err.(*errors.CustomError); ok {
			ctx.JSON(customError.Code, gin.H{
				"error": customError.Message,
				"code":  customError.Code,
			})
		}
	}

	ctx.JSON(http.StatusCreated, comment)
}
*/
