package comments

type CommentsDto struct {
	CourseId string `json:"course_id"`
	UserId   string `json:"user_id"`
	Text     string `json:"text"`
}

type CommentResponse struct {
	Text   string `json:"text"`
	UserId string `json:"id"`
}
type GetCommentsResponse []CommentResponse
