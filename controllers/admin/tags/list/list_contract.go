package list

import (
	"gosimplecms/models"
	"time"
)

type Tag struct {
	ID        uint   `json:"id"`
	Name      string `json:"name,omitempty"`
	Slug      string `json:"slug,omitempty"`
	CreatedAt string `json:"created_at,omitempty"`
}

func (ctl *TagListController) transformToResponse(tags []models.Tag) []*Tag {
	var tagList []*Tag

	if len(tags) == 0 {
		tagList = make([]*Tag, 0)
		return tagList
	}

	for _, tag := range tags {
		tagList = append(tagList, &Tag{
			ID:        tag.ID,
			Name:      tag.Name,
			Slug:      tag.Slug,
			CreatedAt: tag.CreatedAt.Format(time.RFC3339),
		})
	}

	return tagList
}
