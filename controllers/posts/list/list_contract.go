package list

import (
	"gosimplecms/models"
	"time"
)

type Response struct {
	Data []*Post `json:"data"`
}

type Post struct {
	ID          uint   `json:"id"`
	Title       string `json:"title,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Status      string `json:"status,omitempty"`
	PublishedAt string `json:"published_at,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
}

func (ctl *PostListController) transformToResponse(posts []models.Post) []*Post {
	var postList []*Post

	if len(posts) == 0 {
		postList = make([]*Post, 0)
		return postList
	}

	for _, post := range posts {
		postList = append(postList, &Post{
			ID:        post.ID,
			Title:     post.Title,
			Slug:      post.Slug,
			Status:    post.Status,
			CreatedAt: post.CreatedAt.Format(time.RFC3339),
		})
	}

	return postList
}
