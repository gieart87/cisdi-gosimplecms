package list

import (
	"gosimplecms/models"
	"time"
)

type Category struct {
	ID        string `json:"id"`
	Slug      string `json:"slug,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (ctl *CategoryListController) transformToResponse(categories []models.Category) []*Category {
	var categoryList []*Category

	if len(categories) == 0 {
		categoryList = make([]*Category, 0)
		return categoryList
	}

	for _, category := range categories {
		categoryList = append(categoryList, &Category{
			ID:        category.ID,
			Slug:      category.Slug,
			CreatedAt: category.CreatedAt.Format(time.RFC3339),
		})
	}

	return categoryList
}
