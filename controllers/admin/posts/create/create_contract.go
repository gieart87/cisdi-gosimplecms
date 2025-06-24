package create

import (
	"gosimplecms/models"
	"time"
)

func (ctl *PostCreateController) transformToResponse(post *models.Post) *models.CreateUpdatePostResponse {
	return &models.CreateUpdatePostResponse{
		ID:        post.ID,
		Title:     post.Title,
		Status:    post.Status,
		CreatedAt: post.CreatedAt.Format(time.RFC3339),
		UpdatedAt: post.UpdatedAt.Format(time.RFC3339),
	}
}
