package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type PostVersion struct {
	gorm.Model
	Title         string `json:"title"`
	Slug          string `gorm:"size:255;not null;index"`
	VersionNumber int64  `gorm:"not null;index"`
	Status        string `gorm:"size:20;not null;index;default:'DRAFT'"`
	Content       string `json:"content"`
	PostID        uint
	Post          Post
	Categories    []Category `gorm:"many2many:post_version_categories;" json:"categories"`
	Tags          []Tag      `gorm:"many2many:post_version_tags;" json:"tags"`
}

func (p *PostVersion) BeforeCreate(tx *gorm.DB) (err error) {
	p.Slug = slug.Make(p.Title)
	return nil
}
