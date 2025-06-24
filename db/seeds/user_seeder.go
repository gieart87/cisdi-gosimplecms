package seeds

import (
	"fmt"
	"gorm.io/gorm"
	"gosimplecms/models"
	"gosimplecms/utils/password"
)

func SeedUsers(db *gorm.DB) error {
	users := []models.User{
		{Name: "Admin User", Email: "admin@example.com", Role: "ADMIN"},
		{Name: "Editor User", Email: "editor@example.com", Role: "EDITOR"},
		{Name: "Author User", Email: "author@example.com", Role: "AUTHOR"},
		{Name: "Normal User", Email: "user@example.com", Role: "USER"},
	}

	for i := range users {
		pass := password.HashPassword("password")
		users[i].Password = pass
	}

	for _, u := range users {
		if err := db.Where("email = ?", u.Email).FirstOrCreate(&u).Error; err != nil {
			return fmt.Errorf("failed seeding user %s: %w", u.Email, err)
		}
	}
	return nil
}
