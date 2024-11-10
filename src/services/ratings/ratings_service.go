package ratings

import (
	"context"
	"courses-api/src/clients/ratings"
	Ratings_Dto "courses-api/src/dto/ratings"
	"courses-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingsService struct {
	client ratings.RatingsClientInterface
}

func NewRatingsService(client ratings.RatingsClientInterface) RatingsInterface {
	return &RatingsService{
		client: client,
	}
}

type RatingsInterface interface {
	NewRating(ctx context.Context, dto *Ratings_Dto.RatingRequestResponseDto) (*Ratings_Dto.RatingRequestResponseDto, error)
	UpdateRating(ctx context.Context, dto *Ratings_Dto.RatingRequestResponseDto) (*Ratings_Dto.RatingRequestResponseDto, error)
	GetCourseRating(ctx context.Context, courseID string) (*Ratings_Dto.GetCourseRatingResponseDto, error)
}

func (s *RatingsService) NewRating(ctx context.Context, dto *Ratings_Dto.RatingRequestResponseDto) (*Ratings_Dto.RatingRequestResponseDto, error) {
	courseID := dto.CourseId

	rating := models.Rating{
		CourseID: courseID,
		UserID:   dto.UserId,
		Rating:   dto.Rating,
	}

	createdRating, err := s.client.NewRating(ctx, rating)
	if err != nil {
		return nil, err
	}

	return &Ratings_Dto.RatingRequestResponseDto{
		CourseId: createdRating.CourseID,
		UserId:   createdRating.UserID,
		Rating:   createdRating.Rating,
	}, nil
}

func (s *RatingsService) UpdateRating(ctx context.Context, dto *Ratings_Dto.RatingRequestResponseDto) (*Ratings_Dto.RatingRequestResponseDto, error) {
	courseID := dto.CourseId

	rating := models.Rating{
		CourseID: courseID,
		UserID:   dto.UserId,
		Rating:   dto.Rating,
	}

	updatedRating, err := s.client.UpdateRating(ctx, rating)
	if err != nil {
		return nil, err
	}

	return &Ratings_Dto.RatingRequestResponseDto{
		CourseId: updatedRating.CourseID,
		UserId:   updatedRating.UserID,
		Rating:   updatedRating.Rating,
	}, nil
}

func (s *RatingsService) GetCourseRating(ctx context.Context, courseID string) (*Ratings_Dto.GetCourseRatingResponseDto, error) {
	courseObjectID, err := primitive.ObjectIDFromHex(courseID)
	if err != nil {
		return nil, err
	}

	rating, err := s.client.GetRatings(ctx, courseObjectID)
	if err != nil {
		return nil, err
	}

	return &Ratings_Dto.GetCourseRatingResponseDto{
		CourseId: rating.CourseID,
		Rating:   rating.Rating,
	}, nil
}
