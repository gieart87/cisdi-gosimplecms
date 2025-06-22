package models

import "time"

type Tag struct {
	ID           string        `gorm:"size:36;not null;uniqueIndex;primary_key"`
	PostVersions []PostVersion `gorm:"many2many:post_version_tags;"`
	Name         string        `gorm:"size:255;not null;"`
	Slug         string        `gorm:"size:255;not null;uniqueIndex"`
	UsageCount   int           `gorm:"default:0"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
