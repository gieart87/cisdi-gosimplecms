package detail

import (
	"gosimplecms/models"
	"time"
)

type PostResponse struct {
	ID                   uint       `json:"id"`
	Title                string     `json:"title"`
	Status               string     `json:"status"`
	CreatedAt            string     `json:"created_at"`
	UpdatedAt            string     `json:"updated_at"`
	CurrentVersionNumber int64      `json:"current_version_number"`
	Categories           []Category `json:"categories,omitempty"`
	Tags                 []Tag      `json:"tags,omitempty"`
}

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name,omitempty"`
}

type Tag struct {
	ID         uint   `json:"id"`
	Name       string `json:"name,omitempty"`
	UsageCount int64  `gorm:"usage_count;default:0"`
}

func (ctl *PostDetailController) transformToResponse(post *models.Post) *PostResponse {
	return &PostResponse{
		ID:                   post.ID,
		Title:                post.Title,
		Status:               post.Status,
		CreatedAt:            post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:            post.UpdatedAt.Format(time.RFC3339),
		CurrentVersionNumber: post.CurrentVersionNumber,
		Categories: func() []Category {
			var categories []Category
			for _, category := range post.Categories {
				categories = append(categories, Category{
					ID:   category.ID,
					Name: category.Name,
				})
			}
			return categories
		}(),
		Tags: func() []Tag {
			var tags []Tag
			for _, tag := range post.Tags {
				tags = append(tags, Tag{
					ID:         tag.ID,
					Name:       tag.Name,
					UsageCount: tag.UsageCount,
				})
			}

			return tags
		}(),
	}
}
