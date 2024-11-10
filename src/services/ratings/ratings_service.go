package ratings

import (
	"context"
	"courses-api/src/clients"
	Ratings_Dto "courses-api/src/dto/ratings"
	"courses-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RatingsService struct {
	clients *clients.Clients
}

func NewRatingsService(clients *clients.Clients) RatingsInterface {
	return &RatingsService{
		clients: clients,
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

	createdRating, err := s.clients.Ratings.NewRating(ctx, rating)
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

	updatedRating, err := s.clients.Ratings.UpdateRating(ctx, rating)
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

	ratings, err := s.clients.Ratings.GetRatings(ctx, courseObjectID)
	if err != nil {
		return nil, err
	}

	if len(ratings) == 0 {
		return &Ratings_Dto.GetCourseRatingResponseDto{
			CourseId: courseObjectID,
			Rating:   0,
		}, nil
	}

	// Calcula el promedio
	total := 0
	for _, r := range ratings {
		total += r.Rating
	}
	averageRating := total / len(ratings)

	return &Ratings_Dto.GetCourseRatingResponseDto{
		CourseId: courseObjectID,
		Rating:   averageRating,
	}, nil
}
