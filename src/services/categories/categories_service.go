package categories

import (
	"context"
	"courses-api/src/clients"
	dto "courses-api/src/dto/categories"
	"courses-api/src/models"
)

type CategoriesService struct {
	clients *clients.Clients
}

func NewCategoriesService(clients *clients.Clients) CategoryInterface {
	return &CategoriesService{
		clients: clients,
	}
}

type CategoryInterface interface {
	Create(ctx context.Context, dto *dto.CategoryDto) (*dto.CategoryDto, error)
	GetAll(ctx context.Context) (dto.GetCategoriesResponse, error)
}

func (s *CategoriesService) Create(ctx context.Context, dto *dto.CategoryDto) (*dto.CategoryDto, error) {
	category := &models.Category{
		Name: dto.Name,
	}

	result, err := s.clients.Categories.Create(ctx, category)
	if err != nil {
		return nil, err
	}

	dto.Name = result.Name

	return dto, nil
}

func (s *CategoriesService) GetAll(ctx context.Context) (dto.GetCategoriesResponse, error) {
	cats, err := s.clients.Categories.GetAll(ctx)
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
