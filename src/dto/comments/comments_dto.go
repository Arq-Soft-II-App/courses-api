package comments

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentsDto struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CourseId primitive.ObjectID `json:"course_id"`
	UserId   string             `json:"user_id"`
	Text     string             `json:"text"`
}

type CommentResponse struct {
	CourseId primitive.ObjectID `json:"course_id"`
	Text     string             `json:"text"`
	UserId   string             `json:"user_id"`
}

type GetCommentsResponse []CommentResponse
