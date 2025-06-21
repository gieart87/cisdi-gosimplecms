package list

import (
	"gosimplecms/models"
	"time"
)

type Response struct {
	Data []*Post `json:"data"`
}

type Post struct {
	ID          string `json:"id"`
	Title       string `json:"title,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Status      string `json:"status,omitempty"`
	PublishedAt string `json:"published_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func (ctl *PostListController) transformToResponse(posts []models.Post) *Response {
	var response Response

	if len(posts) == 0 {
		response.Data = make([]*Post, 0)
		return &response
	}

	for _, post := range posts {
		response.Data = append(response.Data, &Post{
			ID:     post.ID,
			Title:  post.Title,
			Slug:   post.Slug,
			Status: post.Status,
			PublishedAt: func() string {
				if post.PublishedAt.Time.IsZero() {
					return ""
				}

				return post.PublishedAt.Time.Format(time.RFC3339)
			}(),
			CreatedAt: func() string {
				if post.CreatedAt.IsZero() {
					return ""
				}

				return post.CreatedAt.Format(time.RFC3339)
			}(),
			UpdatedAt: func() string {
				if post.UpdatedAt.IsZero() {
					return ""
				}

				return post.UpdatedAt.Format(time.RFC3339)
			}(),
		})
	}

	return &response
}
