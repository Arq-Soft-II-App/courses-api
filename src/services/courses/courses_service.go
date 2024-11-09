package courses

import (
	"context"
	"courses-api/src/clients/courses"
	dto "courses-api/src/dto/courses"
	"courses-api/src/models"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseInterface interface {
	Create(ctx context.Context, courseDto *dto.CreateCoursesRequestDto) (*dto.CreateCoursesResponseDto, error)
	GetAll(ctx context.Context) (dto.GetAllCourses, error)
	GetById(ctx context.Context, id string) (*dto.GetCourseDto, error)
	Update(ctx context.Context, courseDto *dto.UpdateCourseDto) (*dto.GetCourseDto, error)
	Delete(ctx context.Context, id string) (string, error)
}

type CoursesService struct {
	client courses.CourseClientInterface
}

func NewCoursesService(client courses.CourseClientInterface) CourseInterface {
	return &CoursesService{
		client: client,
	}
}

func (s *CoursesService) Create(ctx context.Context, courseDto *dto.CreateCoursesRequestDto) (*dto.CreateCoursesResponseDto, error) {
	categoryId, err := primitive.ObjectIDFromHex(courseDto.CategoryID)
	if err != nil {
		return nil, err
	}

	if courseDto.CourseImage == "" {
		courseDto.CourseImage = "https://upload.wikimedia.org/wikipedia/commons/a/a3/Image-not-found.png"
	}
	if courseDto.CourseState == nil {
		defaultState := true
		courseDto.CourseState = &defaultState
	}
	if courseDto.CourseCapacity == 0 {
		courseDto.CourseCapacity = 15
	}

	course := models.Course{
		CourseName:        courseDto.CourseName,
		CourseDescription: courseDto.CourseDescription,
		CoursePrice:       courseDto.CoursePrice,
		CourseDuration:    courseDto.CourseDuration,
		CourseInitDate:    courseDto.CourseInitDate,
		CourseState:       *courseDto.CourseState,
		CourseCapacity:    courseDto.CourseCapacity,
		CourseImage:       courseDto.CourseImage,
		CategoryID:        categoryId,
	}

	createdCourse, err := s.client.Create(ctx, course)
	if err != nil {
		return nil, err
	}

	return &dto.CreateCoursesResponseDto{
		CourseName: createdCourse.CourseName,
		CourseId:   createdCourse.Id.Hex(),
	}, nil
}

func (s *CoursesService) GetAll(ctx context.Context) (dto.GetAllCourses, error) {
	courses, err := s.client.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make(dto.GetAllCourses, len(courses))
	for i, course := range courses {
		response[i] = dto.GetCourseDto{
			Id:                 course.Id.Hex(),
			BaseCourseDto:      mapCourseToBaseDto(course),
			CourseCategoryName: course.CategoryName,
			RatingAvg:          course.RatingAvg,
		}
	}
	return response, nil
}

func (s *CoursesService) GetById(ctx context.Context, id string) (*dto.GetCourseDto, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	course, err := s.client.GetById(ctx, objectId)
	if err != nil {
		return nil, err
	}

	return &dto.GetCourseDto{
		Id:                 course.Id.Hex(),
		BaseCourseDto:      mapCourseToBaseDto(*course),
		CourseCategoryName: course.CategoryName,
		RatingAvg:          course.RatingAvg,
	}, nil
}

func (s *CoursesService) Update(ctx context.Context, courseDto *dto.UpdateCourseDto) (*dto.GetCourseDto, error) {
	objectId, err := primitive.ObjectIDFromHex(courseDto.Id)
	if err != nil {
		return nil, err
	}

	originalCourse, err := s.client.GetById(ctx, objectId)
	if err != nil {
		return nil, err
	}

	updatedCourse := models.Course{
		Id:                objectId,
		CourseName:        getUpdatedValue(courseDto.CourseName, originalCourse.CourseName),
		CourseDescription: getUpdatedValue(courseDto.CourseDescription, originalCourse.CourseDescription),
		CoursePrice:       getUpdatedValue(courseDto.CoursePrice, originalCourse.CoursePrice),
		CourseDuration:    getUpdatedValue(courseDto.CourseDuration, originalCourse.CourseDuration),
		CourseInitDate:    getUpdatedValue(courseDto.CourseInitDate, originalCourse.CourseInitDate),
		CourseState:       getUpdatedBoolValue(courseDto.CourseState, originalCourse.CourseState),
		CourseCapacity:    getUpdatedValue(courseDto.CourseCapacity, originalCourse.CourseCapacity),
		CourseImage:       getUpdatedValue(courseDto.CourseImage, originalCourse.CourseImage),
		CategoryID:        originalCourse.CategoryID,
	}

	if courseDto.CategoryID != "" {
		categoryId, err := primitive.ObjectIDFromHex(courseDto.CategoryID)
		if err != nil {
			return nil, err
		}
		updatedCourse.CategoryID = categoryId
	}

	updatedCourseResult, err := s.client.Update(ctx, objectId, updatedCourse)
	if err != nil {
		return nil, err
	}

	return &dto.GetCourseDto{
		Id:                 updatedCourseResult.Id.Hex(),
		BaseCourseDto:      mapCourseToBaseDto(*updatedCourseResult),
		CourseCategoryName: updatedCourseResult.CategoryName,
		RatingAvg:          updatedCourseResult.RatingAvg,
	}, nil
}

func (s *CoursesService) Delete(ctx context.Context, id string) (string, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return "", err
	}

	return s.client.Delete(ctx, objectId)
}

func mapCourseToBaseDto(course models.Course) dto.BaseCourseDto {
	courseState := course.CourseState
	return dto.BaseCourseDto{
		CourseName:        course.CourseName,
		CourseDescription: course.CourseDescription,
		CoursePrice:       course.CoursePrice,
		CourseDuration:    course.CourseDuration,
		CourseInitDate:    course.CourseInitDate,
		CourseState:       &courseState,
		CourseCapacity:    course.CourseCapacity,
		CourseImage:       course.CourseImage,
		CategoryID:        course.CategoryID.Hex(),
	}
}

func getUpdatedValue[T comparable](newValue, originalValue T) T {
	var zero T
	if newValue == zero {
		return originalValue
	}
	return newValue
}
func getUpdatedBoolValue(newValue *bool, originalValue bool) bool {
	if newValue != nil {
		return *newValue
	}
	return originalValue
}
