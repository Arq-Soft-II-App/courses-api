package categories

import (
	"context"
	"courses-api/src/clients/categories"
	dto "courses-api/src/dto/categories"
	"courses-api/src/models"
)

type CategoriesService struct {
	client categories.CategoryClientInterface
}

func NewCategoriesService(client categories.CategoryClientInterface) CategoryInterface {
	return &CategoriesService{
		client: client,
	}
}

type CategoryInterface interface {
	Create(ctx context.Context, dto *dto.CategoryDto) error
	GetAll(ctx context.Context) (dto.GetCategoriesResponse, error)
}

func (s *CategoriesService) Create(ctx context.Context, dto *dto.CategoryDto) error {
	category := &models.Category{
		Name: dto.Name,
	}
	return s.client.Create(ctx, category)
}

func (s *CategoriesService) GetAll(ctx context.Context) (dto.GetCategoriesResponse, error) {
	cats, err := s.client.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	response := make(dto.GetCategoriesResponse, len(cats))
	for i, cat := range cats {
		response[i] = dto.CategoryResponse{
			ID:   cat.ID.Hex(),
			Name: cat.Name,
		}
	}
	return response, nil
}
