package seeds

import (
	"fmt"
	"gosimplecms/models"
	"strings"

	"gorm.io/gorm"
)

func toSlug(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}

func SeedPosts(db *gorm.DB) error {
	// Ambil tag dan kategori dari DB
	var tags []models.Tag
	var categories []models.Category
	var user models.User

	if err := db.Find(&tags).Error; err != nil {
		return fmt.Errorf("failed to get tags: %w", err)
	}
	if err := db.Find(&categories).Error; err != nil {
		return fmt.Errorf("failed to get categories: %w", err)
	}
	if len(tags) < 2 || len(categories) < 1 {
		return fmt.Errorf("minimal 2 tag dan 1 kategori dibutuhkan")
	}

	if err := db.Find(&user).Error; err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	posts := []models.Post{
		{
			Title:      "Getting Started with Golang",
			Slug:       toSlug("Getting Started with Golang"),
			Content:    "This post explains how to start with Golang...",
			Status:     models.PostStatusPublished,
			Tags:       tags[:2],
			Categories: categories[:1],
			UserID:     user.ID,
		},
		{
			Title:      "Understanding Clean Architecture",
			Slug:       toSlug("Understanding Clean Architecture"),
			Content:    "Clean architecture helps you structure code...",
			Status:     models.PostStatusDraft,
			Tags:       tags[1:3],
			Categories: categories[:2],
			UserID:     user.ID,
		},
	}

	for _, post := range posts {
		if err := db.Where("slug = ?", post.Slug).FirstOrCreate(&post).Error; err != nil {
			return fmt.Errorf("failed seeding post: %w", err)
		}
	}

	return nil
}
