package ratings

import (
	"context"
	"courses-api/src/errors"
	"courses-api/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type RatingsClient struct {
	collection *mongo.Collection
}

func NewRatingsClient(db *mongo.Database) *RatingsClient {
	return &RatingsClient{
		collection: db.Collection("ratings"),
	}
}

type RatingsClientInterface interface {
	NewRating(ctx context.Context, rating models.Rating) (models.Rating, error)
	UpdateRating(ctx context.Context, rating models.Rating) (models.Rating, error)
	GetRatings(ctx context.Context) (models.Ratings, error)
}

func (c *RatingsClient) NewRating(ctx context.Context, rating models.Rating) (models.Rating, error) {
	result, err := c.collection.InsertOne(ctx, rating)
	if err != nil {
		return models.Rating{}, errors.NewError("INTERNAL_SERVER_ERROR", "Error al crear el rating", 500)
	}

	rating.ID = result.InsertedID.(primitive.ObjectID)
	return rating, nil
}

func (c *RatingsClient) UpdateRating(ctx context.Context, rating models.Rating) (models.Rating, error) {
	filter := bson.M{"course_id": rating.CourseID, "user_id": rating.UserID}
	update := bson.M{"$set": bson.M{"rating": rating.Rating}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := c.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if result.Err() != nil {
		return models.Rating{}, errors.NewError("INTERNAL_SERVER_ERROR", "Error al actualizar el rating", 500)
	}

	var updatedRating models.Rating
	if err := result.Decode(&updatedRating); err != nil {
		return models.Rating{}, errors.NewError("INTERNAL_SERVER_ERROR", "Error al decodificar rating actualizado", 500)
	}
	return updatedRating, nil
}

func (c *RatingsClient) GetRatings(ctx context.Context) (models.Ratings, error) {
	cursor, err := c.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.NewError("INTERNAL_SERVER_ERROR", "Error al obtener los ratings del curso", 500)
	}
	defer cursor.Close(ctx)

	var ratings models.Ratings
	if err = cursor.All(ctx, &ratings); err != nil {
		return nil, errors.NewError("INTERNAL_SERVER_ERROR", "Error al decodificar ratings", 500)
	}

	if len(ratings) == 0 {
		return nil, errors.NewError("NOT_FOUND", "No se encontraron ratings para el curso", 404)
	}

	return ratings, nil
}
