package courses

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

type CourseClientInterface interface {
	Create(ctx context.Context, course models.Course) (*models.Course, error)
	GetAll(ctx context.Context) (models.Courses, error)
	GetById(ctx context.Context, id primitive.ObjectID) (*models.Course, error)
	Update(ctx context.Context, id primitive.ObjectID, course models.Course) (*models.Course, error)
	Delete(ctx context.Context, id primitive.ObjectID) (string, error)
}

type CourseClient struct {
	collection *mongo.Collection
}

func NewCourseClient(db *mongo.Database) *CourseClient {
	return &CourseClient{
		collection: db.Collection("courses"),
	}
}

func (c *CourseClient) Create(ctx context.Context, course models.Course) (*models.Course, error) {
	result, err := c.collection.InsertOne(ctx, course)
	if err != nil {
		return nil, errors.NewError("COURSE_CREATION_FAILED", fmt.Sprintf("Error al crear el curso: %v", err), 500)
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.NewError("INVALID_OBJECT_ID", "Error al obtener el ID insertado", 500)
	}
	course.Id = oid

	return &course, nil
}

func (c *CourseClient) GetAll(ctx context.Context) (models.Courses, error) {
	pipeline := c.buildCoursePipeline(nil)
	cursor, err := c.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.NewError("COURSE_FETCH_FAILED", fmt.Sprintf("Error al obtener cursos: %v", err), 500)
	}
	defer cursor.Close(ctx)

	var courses models.Courses
	if err := cursor.All(ctx, &courses); err != nil {
		return nil, errors.NewError("COURSE_DECODE_FAILED", fmt.Sprintf("Error al decodificar cursos: %v", err), 500)
	}

	return courses, nil
}

func (c *CourseClient) GetById(ctx context.Context, id primitive.ObjectID) (*models.Course, error) {
	matchStage := bson.D{{Key: "$match", Value: bson.M{"_id": id}}}
	pipeline := c.buildCoursePipeline(matchStage)
	cursor, err := c.collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, errors.NewError("COURSE_FETCH_FAILED", fmt.Sprintf("Error al obtener el curso por ID: %v", err), 500)
	}
	defer cursor.Close(ctx)

	if !cursor.Next(ctx) {
		return nil, errors.NewError("COURSE_NOT_FOUND", "Curso no encontrado", 404)
	}

	var course models.Course
	if err := cursor.Decode(&course); err != nil {
		return nil, errors.NewError("COURSE_DECODE_FAILED", fmt.Sprintf("Error al decodificar el curso: %v", err), 500)
	}

	return &course, nil
}

func (c *CourseClient) Update(ctx context.Context, id primitive.ObjectID, course models.Course) (*models.Course, error) {
	update := bson.M{
		"$set": bson.M{
			"course_name": course.CourseName,
			"description": course.CourseDescription,
			"price":       course.CoursePrice,
			"duration":    course.CourseDuration,
			"init_date":   course.CourseInitDate,
			"state":       course.CourseState,
			"cupo":        course.CourseCapacity,
			"image":       course.CourseImage,
			"category_id": course.CategoryID,
		},
	}

	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)
	result := c.collection.FindOneAndUpdate(ctx, bson.M{"_id": id}, update, opts)
	if err := result.Err(); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewError("COURSE_NOT_FOUND", "Curso no encontrado", 404)
		}
		return nil, errors.NewError("COURSE_UPDATE_FAILED", fmt.Sprintf("Error al actualizar el curso: %v", err), 500)
	}

	var updatedCourse models.Course
	if err := result.Decode(&updatedCourse); err != nil {
		return nil, errors.NewError("COURSE_DECODE_FAILED", fmt.Sprintf("Error al decodificar el curso actualizado: %v", err), 500)
	}

	return &updatedCourse, nil
}

func (c *CourseClient) Delete(ctx context.Context, id primitive.ObjectID) (string, error) {
	_, err := c.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return "", errors.NewError("COURSE_DELETE_FAILED", fmt.Sprintf("Error al eliminar el curso: %v", err), 500)
	}
	return fmt.Sprintf("Curso %s eliminado", id.Hex()), nil
}

// buildCoursePipeline construye el pipeline de agregación para obtener cursos.
func (c *CourseClient) buildCoursePipeline(matchStage interface{}) mongo.Pipeline {
	pipeline := mongo.Pipeline{}

	if matchStage != nil {
		pipeline = append(pipeline, matchStage.(bson.D))
	}

	pipeline = append(pipeline,
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "categories",
			"localField":   "category_id",
			"foreignField": "_id",
			"as":           "category",
		}}},
		bson.D{{Key: "$unwind", Value: "$category"}},
		bson.D{{Key: "$lookup", Value: bson.M{
			"from":         "ratings",
			"localField":   "_id",
			"foreignField": "course_id",
			"as":           "ratings",
		}}},
		bson.D{{Key: "$addFields", Value: bson.M{
			"ratingavg": bson.M{
				"$cond": bson.M{
					"if":   bson.M{"$eq": bson.A{"$ratings", bson.A{}}},
					"then": 0,
					"else": bson.M{
						"$avg": "$ratings.rating",
					},
				},
			},
		}}},
		bson.D{{Key: "$project", Value: bson.M{
			"_id":           1,
			"course_name":   1,
			"description":   1,
			"price":         1,
			"duration":      1,
			"init_date":     1,
			"state":         1,
			"cupo":          1,
			"image":         1,
			"category_id":   1,
			"category_name": "$category.name",
			"ratingavg":     1,
		}}},
	)

	return pipeline
}
