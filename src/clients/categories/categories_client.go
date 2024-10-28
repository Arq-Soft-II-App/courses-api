package categories

import (
	"context"
	"courses-api/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoriesClient struct {
	collection *mongo.Collection
}

func NewCategoriesClient(db *mongo.Database) *CategoriesClient {
	return &CategoriesClient{
		collection: db.Collection("categories"),
	}
}

func (c *CategoriesClient) Create(ctx context.Context, category *models.Category) error {
	_, err := c.collection.InsertOne(ctx, category)
	return err
}

func (c *CategoriesClient) GetAll(ctx context.Context) ([]models.Category, error) {
	cursor, err := c.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var categories []models.Category
	if err := cursor.All(ctx, &categories); err != nil {
		return nil, err
	}
	return categories, nil
}
