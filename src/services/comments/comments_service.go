package comments

import (
	"context"
	"courses-api/src/clients/comments"
	Comments_Dto "courses-api/src/dto/comments"
	"courses-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentsService struct {
	client comments.CommentsClientInterface
}

func NewCommentsService(client comments.CommentsClientInterface) CommentsInterface {
	return &CommentsService{
		client: client,
	}
}

type CommentsInterface interface {
	NewComment(ctx context.Context, dto *Comments_Dto.CommentsDto) (*Comments_Dto.CommentResponse, error)
	GetCourseComments(ctx context.Context, courseID string) (Comments_Dto.GetCommentsResponse, error)
	UpdateComment(ctx context.Context, dto *Comments_Dto.CommentsDto) (*Comments_Dto.CommentResponse, error)
}

func (s *CommentsService) NewComment(ctx context.Context, dto *Comments_Dto.CommentsDto) (*Comments_Dto.CommentResponse, error) {
	courseID, err := primitive.ObjectIDFromHex(dto.CourseId.Hex())
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

	return &Comments_Dto.CommentResponse{
		Text:   createdComment.Text,
		UserId: createdComment.UserId,
	}, nil
}

func (s *CommentsService) GetCourseComments(ctx context.Context, courseID string) (Comments_Dto.GetCommentsResponse, error) {
	courseObjectID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, err
	}

	comments, err := s.client.GetCourseComments(ctx, courseObjectID)
	if err != nil {
		return nil, err
	}

	response := make(Comments_Dto.GetCommentsResponse, len(comments))
	for i, comment := range comments {
		response[i] = Comments_Dto.CommentResponse{
			Text:   comment.Text,
			UserId: comment.UserId,
		}
	}

	return response, nil
}

func (s *CommentsService) UpdateComment(ctx context.Context, dto *Comments_Dto.CommentsDto) (*Comments_Dto.CommentResponse, error) {
	CourseId, err := primitive.ObjectIDFromHex(dto.CourseId.Hex())
	if err != nil {
		return nil, err
	}

	comment := models.Comment{
		Text:     dto.Text,
		UserId:   dto.UserId,
		CourseId: CourseId,
	}

	updatedComment, err := s.client.UpdateComment(ctx, comment)
	if err != nil {
		return nil, err
	}

	return &Comments_Dto.CommentResponse{
		Text:   updatedComment.Text,
		UserId: updatedComment.UserId,
	}, nil
}
