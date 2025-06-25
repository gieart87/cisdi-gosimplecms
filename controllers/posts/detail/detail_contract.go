package detail

import (
	"gosimplecms/models"
	"log"
	"time"
)

type PostResponse struct {
	ID                   uint         `json:"id"`
	Title                string       `json:"title"`
	Status               string       `json:"status"`
	CreatedAt            string       `json:"created_at"`
	UpdatedAt            string       `json:"updated_at"`
	CurrentVersionNumber int64        `json:"current_version_number"`
	Categories           []Category   `json:"categories,omitempty"`
	Tags                 []Tag        `json:"tags,omitempty"`
	TagScorePost         TagScorePost `json:"tag_score_post"`
}

type TagScorePost struct {
	Scores     []TagScore `json:"scores"`
	TotalScore float64    `json:"total_score"`
}
type TagScore struct {
	Tag1ID   uint    `json:"tag_1_id"`
	Tag2ID   uint    `json:"tag_2_id"`
	Tag1Name string  `json:"tag_1_name"`
	Tag2Name string  `json:"tag_2_name"`
	Score    float64 `json:"score"`
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

var PostTagIDs []uint

func (ctl *PostDetailController) getTagName(tagId uint) string {
	tags, err := ctl.PostService.FindTagsByIDs(PostTagIDs)
	if err != nil {
		return ""
	}

	for _, tag := range tags {
		if tag.ID == tagId {
			return tag.Name
		}
	}

	return ""
}

func (ctl *PostDetailController) transformToResponse(post *models.Post) *PostResponse {
	var tagScores []TagScore

	var tagIDs []uint
	for _, tag := range post.Tags {
		tagIDs = append(tagIDs, tag.ID)
	}

	PostTagIDs = tagIDs

	pairScores, totalScore, err := ctl.PostService.CalculateTagRelationshipScore(tagIDs)
	if err != nil {
		log.Fatal(err)
	}

	for _, pair := range pairScores {
		tagScores = append(tagScores, TagScore{
			Tag1ID:   pair.Tag1ID,
			Tag2ID:   pair.Tag2ID,
			Score:    pair.Score,
			Tag1Name: ctl.getTagName(pair.Tag1ID),
			Tag2Name: ctl.getTagName(pair.Tag2ID),
		})
	}

	var tagScorePost TagScorePost
	tagScorePost.Scores = tagScores
	tagScorePost.TotalScore = totalScore

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
		TagScorePost: tagScorePost,
	}
}
