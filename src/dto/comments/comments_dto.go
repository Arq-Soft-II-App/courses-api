package comments

import "go.mongodb.org/mongo-driver/bson/primitive"

// CommentsDto representa la estructura para la creación de comentarios
type CommentsDto struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CourseId string             `json:"course_id"`
	UserId   string             `json:"user_id"`
	Text     string             `json:"text"`
}

// CommentResponse representa la respuesta de un comentario
type CommentResponse struct {
	Text   string `json:"text"`
	UserId string `json:"user_id"`
}

// GetCommentsResponse representa una lista de respuestas de comentarios
type GetCommentsResponse []CommentResponse
