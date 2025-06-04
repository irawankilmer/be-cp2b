package mapper

import (
	"be-cp2b/internal/domain"
	"be-cp2b/internal/dto/response"
)

func MapCategoryToDTO(t domain.Category) response.CategoryResponse {
	return response.CategoryResponse{
		ID:           t.ID,
		Name:         t.Name,
		Type:         t.Type,
		Descriptions: t.Descriptions,
	}
}
