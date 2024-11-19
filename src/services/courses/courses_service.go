package courses

import (
	"context"
	"courses-api/src/clients"
	rabbitmq "courses-api/src/config/rabbitMQ"
	dto "courses-api/src/dto/courses"
	"courses-api/src/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CourseInterface interface {
	Create(ctx context.Context, courseDto *dto.CreateCoursesRequestDto) (*dto.CreateCoursesResponseDto, error)
	GetAll(ctx context.Context) (dto.GetAllCourses, error)
	GetById(ctx context.Context, id string) (*dto.GetCourseDto, error)
	Update(ctx context.Context, courseDto *dto.UpdateCourseDto) (*dto.GetCourseDto, error)
	Delete(ctx context.Context, id string) (string, error)
	GetCourseList(ctx context.Context, ids []string) ([]dto.CourseListDto, error)
}

type CoursesService struct {
	clients  *clients.Clients
	rabbitMQ *rabbitmq.RabbitMQ
}

func NewCoursesService(clients *clients.Clients, rabbitMQ *rabbitmq.RabbitMQ) CourseInterface {
	return &CoursesService{
		clients:  clients,
		rabbitMQ: rabbitMQ,
	}
}

func (s *CoursesService) GetCourseList(ctx context.Context, ids []string) ([]dto.CourseListDto, error) {
	fmt.Printf("Service - Buscando cursos con IDs: %v\n", ids)

	courses, err := s.clients.Courses.GetCourseList(ctx, ids)
	if err != nil {
		fmt.Printf("Error en client.GetCourseList: %v\n", err)
		return nil, err
	}

	var courseResponses []dto.CourseListDto
	for _, course := range courses {
		courseResponses = append(courseResponses, dto.CourseListDto{
			Id:           course.Id.Hex(),
			CategoryID:   course.CategoryID.Hex(),
			CourseName:   course.CourseName,
			Description:  course.CourseDescription,
			Price:        course.CoursePrice,
			Duration:     course.CourseDuration,
			Capacity:     course.CourseCapacity,
			InitDate:     course.CourseInitDate,
			State:        course.CourseState,
			Image:        course.CourseImage,
			CategoryName: course.CategoryName,
			RatingAvg:    course.RatingAvg,
		})
	}

	return courseResponses, nil
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

	createdCourse, err := s.clients.Courses.Create(ctx, course)
	if err != nil {
		return nil, err
	}

	// Publicar mensaje en RabbitMQ
	err = s.rabbitMQ.PublishMessage(createdCourse.Id.Hex())
	if err != nil {
		log.Printf("Error al publicar mensaje en RabbitMQ: %v", err)
	}

	return &dto.CreateCoursesResponseDto{
		CourseName: createdCourse.CourseName,
		CourseId:   createdCourse.Id.Hex(),
	}, nil
}

func (s *CoursesService) GetAll(ctx context.Context) (dto.GetAllCourses, error) {
	courses, err := s.clients.Courses.GetAll(ctx)
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

	course, err := s.clients.Courses.GetById(ctx, objectId)
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

	originalCourse, err := s.clients.Courses.GetById(ctx, objectId)
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

	updatedCourseResult, err := s.clients.Courses.Update(ctx, objectId, updatedCourse)
	if err != nil {
		return nil, err
	}

	// Publicar mensaje en RabbitMQ
	err = s.rabbitMQ.PublishMessage(updatedCourseResult.Id.Hex())
	if err != nil {
		log.Printf("Error al publicar mensaje en RabbitMQ: %v", err)
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

	result, err := s.clients.Courses.Delete(ctx, objectId)
	if err != nil {
		return "", err
	}

	// Publicar mensaje en RabbitMQ
	err = s.rabbitMQ.PublishMessage(id)
	if err != nil {
		log.Printf("Error al publicar mensaje en RabbitMQ: %v", err)
	}

	return result, nil
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
