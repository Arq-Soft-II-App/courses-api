package comments

import "go.mongodb.org/mongo-driver/bson/primitive"

type CommentsDto struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	CourseId string             `json:"course_id"`
	UserId   string             `json:"user_id"`
	Text     string             `json:"text"`
}

type CommentResponse struct {
	Text       string `json:"text"`
	UserId     string `json:"user_id"`
	UserName   string `json:"user_name"`
	UserAvatar string `json:"user_avatar"`
}

type GetCommentsResponse []CommentResponse
