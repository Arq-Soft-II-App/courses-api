package courses

type BaseCourseDto struct {
	CourseName        string  `json:"course_name"`
	CourseDescription string  `json:"description"`
	CoursePrice       float64 `json:"price"`
	CourseDuration    int     `json:"duration"`
	CourseCapacity    int     `json:"capacity"`
	CategoryID        string  `json:"category_id"`
	CourseInitDate    string  `json:"init_date"`
	CourseState       bool    `json:"state"`
	CourseImage       string  `json:"image"`
}

type CreateCoursesRequestDto BaseCourseDto

type CreateCoursesResponseDto struct {
	CourseName string `json:"course_name"`
	CourseId   string `json:"_id"`
}

type GetOneCourseRequestDto struct {
	CourseId string `json:"_id"`
}

type GetOneCourseResponseDto BaseCourseDto

type GetCourseDto struct {
	Id string `json:"_id"`
	BaseCourseDto
	CourseCategoryName string  `json:"category_name"`
	RatingAvg          float64 `json:"ratingavg"`
}

type GetAllCourses []GetCourseDto

type UpdateCourseDto struct {
	Id string `json:"_id"`
	BaseCourseDto
}
