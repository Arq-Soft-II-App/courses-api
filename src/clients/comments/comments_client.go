package comments

import (
	"context"
	"courses-api/src/errors"
	"courses-api/src/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CommentsClient struct {
	collection *mongo.Collection
}

type CommentsClientInterface interface {
	NewComment(ctx context.Context, comment *models.Comment) (models.Comment, error)
	GetCourseComments(ctx context.Context, courseID primitive.ObjectID) ([]models.Comment, error)
	UpdateComment(ctx context.Context, comment models.Comment) (*models.Comment, error)
}

func NewCommentsClient(db *mongo.Database) CommentsClientInterface {
	return &CommentsClient{
		collection: db.Collection("comments"),
	}
}

func (c *CommentsClient) NewComment(ctx context.Context, comment *models.Comment) (models.Comment, error) {
	result, err := c.collection.InsertOne(ctx, comment)
	if err != nil {
		return models.Comment{}, errors.NewError("COMMENT_CREATION_FAILED", fmt.Sprintf("Error al crear el comentario: %v", err), 500)
	}

	comment.ID = result.InsertedID.(primitive.ObjectID)
	return *comment, nil
}

func (c *CommentsClient) GetCourseComments(ctx context.Context, courseID primitive.ObjectID) ([]models.Comment, error) {
	filter := bson.M{"course_id": courseID}
	cursor, err := c.collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.NewError("COMMENT_FETCH_FAILED", fmt.Sprintf("Error al obtener los comentarios: %v", err), 500)
	}
	defer cursor.Close(ctx)

	var comments []models.Comment
	if err = cursor.All(ctx, &comments); err != nil {
		return nil, errors.NewError("COMMENT_DECODE_FAILED", fmt.Sprintf("Error al decodificar comentarios: %v", err), 500)
	}

	if len(comments) == 0 {
		return nil, errors.NewError("COMMENT_NOT_FOUND", "No se encontraron comentarios para el curso", 404)
	}

	return comments, nil
}

func (c *CommentsClient) UpdateComment(ctx context.Context, comment models.Comment) (*models.Comment, error) {
	filter := bson.M{"course_id": comment.CourseId, "user_id": comment.UserId}
	update := bson.M{"$set": bson.M{"text": comment.Text}}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := c.collection.FindOneAndUpdate(ctx, filter, update, opts)
	if err := result.Err(); err != nil {
		return nil, errors.NewError("COMMENT_UPDATE_FAILED", fmt.Sprintf("Error al actualizar el comentario: %v", err), 500)
	}

	var updatedComment models.Comment
	if err := result.Decode(&updatedComment); err != nil {
		return nil, errors.NewError("COMMENT_DECODE_FAILED", fmt.Sprintf("Error al decodificar el comentario actualizado: %v", err), 500)
	}
	return &updatedComment, nil
}
