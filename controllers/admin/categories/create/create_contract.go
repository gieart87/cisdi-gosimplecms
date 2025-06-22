package create

import (
	"gosimplecms/models"
)

func (ctl *CategoryCreateController) transformToResponse(category *models.Category) *models.CreateUpdateCategoryResponse {
	return &models.CreateUpdateCategoryResponse{
		ID:   category.ID,
		Name: category.Name,
		Slug: category.Slug,
	}
}
