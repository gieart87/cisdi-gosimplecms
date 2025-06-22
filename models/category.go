package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Category struct {
	ID           string        `gorm:"size:36;not null;uniqueIndex;primary_key"`
	PostVersions []PostVersion `gorm:"many2many:post_version_categories;"`
	Name         string        `gorm:"size:255;not null;"`
	Slug         string        `gorm:"size:255;not null;uniqueIndex"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *Category) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}
