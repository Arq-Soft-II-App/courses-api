package comments

import (
	"context"
	"courses-api/src/clients/comments"
	dto "courses-api/src/dto/comments"
	"courses-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentsService define el servicio para manejar comentarios
type CommentsService struct {
	client      comments.CommentsClientInterface
	usersClient UsersClientInterface // Interfaz para obtener información del usuario
}

// NewCommentsService inicializa un nuevo servicio de comentarios
func NewCommentsService(client comments.CommentsClientInterface, usersClient UsersClientInterface) CommentsInterface {
	return &CommentsService{
		client:      client,
		usersClient: usersClient,
	}
}

// CommentsInterface define la interfaz para el servicio de comentarios
type CommentsInterface interface {
	NewComment(ctx context.Context, dto *dto.CommentsDto) (*dto.CommentResponse, error)
	GetCourseComments(ctx context.Context, courseID string) (dto.GetCommentsResponse, error)
	UpdateComment(ctx context.Context, dto *dto.CommentsDto) (*dto.CommentResponse, error)
}

// NewComment crea un nuevo comentario
func (s *CommentsService) NewComment(ctx context.Context, dto *dto.CommentsDto) (*dto.CommentResponse, error) {
	courseID, err := primitive.ObjectIDFromHex(dto.CourseId)
	if err != nil {
		return nil, err
	}

	comment := &models.Comment{
		Text:     dto.Text,
		CourseId: courseID,
		UserId:   dto.UserId,
	}

	createdComment, err := s.client.NewComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &dto.CommentResponse{
		Text:   createdComment.Text,
		UserId: createdComment.UserId,
	}, nil
}

// GetCourseComments obtiene los comentarios de un curso específico
func (s *CommentsService) GetCourseComments(ctx context.Context, courseID string) (dto.GetCommentsResponse, error) {
	courseObjectID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, err
	}

	comments, err := s.client.GetCourseComments(ctx, courseObjectID)
	if err != nil {
		return nil, err
	}

	response := make(dto.GetCommentsResponse, len(comments))
	for i, comment := range comments {
		user, err := s.usersClient.GetUser(ctx, comment.UserId) // Obtener datos del usuario
		if err != nil {
			return nil, err
		}

		response[i] = dto.CommentResponse{
			Text:       comment.Text,
			UserId:     comment.UserId,
			UserName:   user.Name,
			UserAvatar: user.Avatar,
		}
	}

	return response, nil
}

// UpdateComment actualiza un comentario existente
func (s *CommentsService) UpdateComment(ctx context.Context, dto *dto.CommentsDto) (*dto.CommentResponse, error) {
	commentID, err := primitive.ObjectIDFromHex(dto.ID.Hex())
	if err != nil {
		return nil, err
	}

	comment := models.Comment{
		ID:       commentID,
		Text:     dto.Text,
		UserId:   dto.UserId,
		CourseId: commentID,
	}

	updatedComment, err := s.client.UpdateComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &dto.CommentResponse{
		Text:   updatedComment.Text,
		UserId: updatedComment.UserId,
	}, nil
}
