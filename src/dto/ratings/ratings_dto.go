package rating

import "go.mongodb.org/mongo-driver/bson/primitive"

type RatingRequestResponseDto struct {
	CourseId primitive.ObjectID `json:"course_id"`
	UserId   string             `json:"user_id"`
	Rating   int                `json:"rating"`
}

type GetCourseRatingRequestDto struct {
	CourseId primitive.ObjectID `json:"course_id"`
}

type GetCourseRatingResponseDto struct {
	CourseId primitive.ObjectID `json:"course_id"`
	Rating   int                `json:"rating"`
}
type CourseRatingsDto []GetCourseRatingResponseDto
type RatingsResponse []RatingRequestResponseDto
