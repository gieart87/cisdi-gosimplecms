package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title            string     `json:"title"`
	Slug             string     `gorm:"size:255;not null;uniqueIndex"`
	Status           string     `gorm:"size:20;not null;index;default:'DRAFT'"`
	CurrentVersionID uint       `json:"current_version_id"`
	Content          string     `json:"content"`
	Categories       []Category `gorm:"many2many:post_categories" json:"categories"`
	Tags             []Tag      `gorm:"many2many:post_tags" json:"tags"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.Slug = slug.Make(p.Title)
	return nil
}

type CreatePostRequest struct {
	Slug        string `json:"slug" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Content     string `json:"content"`
	TagIDs      []uint `json:"tag_ids"`
	CategoryIDs []uint `json:"category_ids"`
	UserID      uint   `json:"user_id"`
}

type PaginationMeta struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

type CreateUpdatePostResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (r CreatePostRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Title, validation.Required, validation.RuneLength(3, 100)),
	)
}
