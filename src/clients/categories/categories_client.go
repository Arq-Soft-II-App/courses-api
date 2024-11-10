package categories

import (
	"context"
	"courses-api/src/errors"
	"courses-api/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoriesClient struct {
	collection *mongo.Collection
}

type CategoryClientInterface interface {
	Create(ctx context.Context, category *models.Category) (models.Category, error)
	GetAll(ctx context.Context) ([]models.Category, error)
}

func NewCategoriesClient(db *mongo.Database) CategoryClientInterface {
	return &CategoriesClient{
		collection: db.Collection("categories"),
	}
}

func (c *CategoriesClient) Create(ctx context.Context, category *models.Category) (models.Category, error) {
	result, err := c.collection.InsertOne(ctx, category)
	if err != nil {
		return models.Category{}, errors.NewError("INTERNAL_SERVER_ERROR", "Error al insertar la categoría en la base de datos", 500)
	}
	category.ID = result.InsertedID.(primitive.ObjectID)
	return *category, nil
}

func (c *CategoriesClient) GetAll(ctx context.Context) ([]models.Category, error) {
	cursor, err := c.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.NewError("INTERNAL_SERVER_ERROR", "Error al obtener las categorías de la base de datos", 500)
	}
	defer cursor.Close(ctx)

	var categories []models.Category
	if err := cursor.All(ctx, &categories); err != nil {
		return nil, errors.NewError("INTERNAL_SERVER_ERROR", "Error al decodificar las categorías de la base de datos", 500)
	}
	return categories, nil
}
