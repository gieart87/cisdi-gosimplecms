package seeds

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"gosimplecms/models"
)

func SeedTagsAndCategories(db *gorm.DB) error {
	tags := []string{
		"Golang", "Backend", "Docker", "API", "Clean Architecture",
	}

	categories := []string{
		"Technology", "Programming", "Tutorial", "DevOps",
	}

	// Seed Tags
	for _, name := range tags {
		tag := models.Tag{
			Name: name,
			Slug: slug.Make(name),
		}
		if err := db.
			Where("name = ?", tag.Name).
			FirstOrCreate(&tag).Error; err != nil {
			return err
		}
	}

	// Seed Categories
	for _, name := range categories {
		category := models.Category{
			Name: name,
			Slug: slug.Make(name),
		}
		if err := db.
			Where("name = ?", category.Name).
			FirstOrCreate(&category).Error; err != nil {
			return err
		}
	}

	return nil
}
