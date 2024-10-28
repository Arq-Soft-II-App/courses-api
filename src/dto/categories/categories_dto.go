package categories

type CategoryDto struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   string `json:"_id"`
	Name string `json:"name"`
}

type GetCategoriesResponse []CategoryResponse
