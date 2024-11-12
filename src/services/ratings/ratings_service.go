package ratings

import (
	"context"
	"courses-api/src/clients"
	rabbitmq "courses-api/src/config/rabbitMQ"
	Ratings_Dto "courses-api/src/dto/ratings"
	"courses-api/src/models"
	"log"
)

type RatingsService struct {
	clients  *clients.Clients
	rabbitMQ *rabbitmq.RabbitMQ
}

func NewRatingsService(clients *clients.Clients, rabbitMQ *rabbitmq.RabbitMQ) RatingsInterface {
	return &RatingsService{
		clients:  clients,
		rabbitMQ: rabbitMQ,
	}
}

type RatingsInterface interface {
	NewRating(ctx context.Context, dto *Ratings_Dto.RatingRequestResponseDto) (*Ratings_Dto.RatingRequestResponseDto, error)
	UpdateRating(ctx context.Context, dto *Ratings_Dto.RatingRequestResponseDto) (*Ratings_Dto.RatingRequestResponseDto, error)
	GetAllRatings(ctx context.Context) (*Ratings_Dto.RatingsResponse, error)
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

	// Publicar mensaje en RabbitMQ
	err = s.rabbitMQ.PublishMessage(createdRating.CourseID.Hex())
	if err != nil {
		log.Printf("Error al publicar mensaje en RabbitMQ: %v", err)
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

	// Publicar mensaje en RabbitMQ
	err = s.rabbitMQ.PublishMessage(updatedRating.CourseID.Hex())
	if err != nil {
		log.Printf("Error al publicar mensaje en RabbitMQ: %v", err)
	}

	return &Ratings_Dto.RatingRequestResponseDto{
		CourseId: updatedRating.CourseID,
		UserId:   updatedRating.UserID,
		Rating:   updatedRating.Rating,
	}, nil
}

func (s *RatingsService) GetAllRatings(ctx context.Context) (*Ratings_Dto.RatingsResponse, error) {
	ratings, err := s.clients.Ratings.GetRatings(ctx)
	if err != nil {
		return nil, err
	}

	var response Ratings_Dto.RatingsResponse
	for _, r := range ratings {
		response = append(response, Ratings_Dto.RatingRequestResponseDto{
			CourseId: r.CourseID,
			UserId:   r.UserID,
			Rating:   r.Rating,
		})
	}

	return &response, nil
}
