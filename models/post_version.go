package models

import "database/sql"

type PostVersion struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Post        Post
	PostID      string       `gorm:"size:36;index"`
	VersionID   string       `gorm:"size:255;uniqueIndex"`
	Categories  []Category   `gorm:"many2many:post_version_categories;"`
	Tags        []Tag        `gorm:"many2many:post_version_tags;"`
	Title       string       `gorm:"size:255"`
	Slug        string       `gorm:"size:255;uniqueIndex"`
	Body        string       `gorm:"type:text"`
	Status      string       `gorm:"size:20;not null;index;default:'DRAFT'"`
	PublishedAt sql.NullTime `gorm:"index"`
}
