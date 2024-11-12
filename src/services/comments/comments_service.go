package comments

import (
	"context"
	"courses-api/src/clients"
	rabbitmq "courses-api/src/config/rabbitMQ"
	Comments_Dto "courses-api/src/dto/comments"
	"courses-api/src/models"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentsService struct {
	clients  *clients.Clients
	rabbitMQ *rabbitmq.RabbitMQ
}

func NewCommentsService(clients *clients.Clients, rabbitMQ *rabbitmq.RabbitMQ) CommentsInterface {
	return &CommentsService{
		clients:  clients,
		rabbitMQ: rabbitMQ,
	}
}

type CommentsInterface interface {
	NewComment(ctx context.Context, dto *Comments_Dto.CommentsDto) (*Comments_Dto.CommentResponse, error)
	GetCourseComments(ctx context.Context, courseID string) (Comments_Dto.GetCommentsResponse, error)
	UpdateComment(ctx context.Context, dto *Comments_Dto.CommentsDto) (*Comments_Dto.CommentResponse, error)
}

func (s *CommentsService) NewComment(ctx context.Context, dto *Comments_Dto.CommentsDto) (*Comments_Dto.CommentResponse, error) {
	courseID := dto.CourseId

	comment := &models.Comment{
		Text:     dto.Text,
		CourseId: courseID,
		UserId:   dto.UserId,
	}

	createdComment, err := s.clients.Comments.NewComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	err = s.rabbitMQ.PublishMessage(createdComment.CourseId.Hex())
	if err != nil {
		log.Printf("Error al publicar mensaje en RabbitMQ: %v", err)
	}

	return &Comments_Dto.CommentResponse{
		CourseId: createdComment.CourseId,
		Text:     createdComment.Text,
		UserId:   createdComment.UserId,
	}, nil
}

func (s *CommentsService) GetCourseComments(ctx context.Context, courseID string) (Comments_Dto.GetCommentsResponse, error) {
	courseObjectID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, err
	}

	comments, err := s.clients.Comments.GetCourseComments(ctx, courseObjectID)
	if err != nil {
		return nil, err
	}

	response := make(Comments_Dto.GetCommentsResponse, len(comments))
	for i, comment := range comments {
		response[i] = Comments_Dto.CommentResponse{
			CourseId: comment.CourseId,
			Text:     comment.Text,
			UserId:   comment.UserId,
		}
	}

	return response, nil
}

func (s *CommentsService) UpdateComment(ctx context.Context, dto *Comments_Dto.CommentsDto) (*Comments_Dto.CommentResponse, error) {
	courseID := dto.CourseId

	comment := models.Comment{
		Text:     dto.Text,
		UserId:   dto.UserId,
		CourseId: courseID,
	}

	updatedComment, err := s.clients.Comments.UpdateComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &Comments_Dto.CommentResponse{
		CourseId: updatedComment.CourseId,
		Text:     updatedComment.Text,
		UserId:   updatedComment.UserId,
	}, nil
}
