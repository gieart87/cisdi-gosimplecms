package create

import (
	"gosimplecms/models"
)

func (ctl *TagCreateController) transformToResponse(category *models.Tag) *models.CreateUpdateTagResponse {
	return &models.CreateUpdateTagResponse{
		ID:   category.ID,
		Name: category.Name,
	}
}
