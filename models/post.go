package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"regexp"
)

type Post struct {
	gorm.Model
	Title                string `json:"title"`
	Slug                 string `gorm:"size:255;not null;uniqueIndex"`
	Status               string `gorm:"size:20;not null;index;default:'DRAFT'"`
	CurrentVersionNumber int64  `json:"current_version_number"`
	Content              string `json:"content"`
	UserID               uint   `json:"user_id"`
	User                 User
	Categories           []Category `gorm:"many2many:post_categories" json:"categories"`
	Tags                 []Tag      `gorm:"many2many:post_tags" json:"tags"`
}

const (
	PostStatusPublished = "PUBLISHED"
	PostStatusDraft     = "DRAFT"
	PostStatusArchived  = "ARCHIVED"
)

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.Slug = slug.Make(p.Title)
	return nil
}

type CreatePostRequest struct {
	Slug        string `json:"slug"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	TagIDs      []uint `json:"tag_ids"`
	CategoryIDs []uint `json:"category_ids"`
	UserID      uint   `json:"user_id"`
}

type UpdatePostRequest struct {
	Title       string `json:"title"`
	Content     string `json:"content"`
	TagIDs      []uint `json:"tag_ids"`
	CategoryIDs []uint `json:"category_ids"`
}

type PaginationMeta struct {
	Total int64 `json:"total"`
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
}

type CreateUpdatePostResponse struct {
	ID                   uint   `json:"id"`
	Title                string `json:"title"`
	Status               string `json:"status"`
	CreatedAt            string `json:"created_at"`
	UpdatedAt            string `json:"updated_at"`
	CurrentVersionNumber int64  `json:"current_version_number"`
}

func (r CreatePostRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Title,
			validation.Required.Error("title is required"),
			validation.Match(regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)).Error("only letters, numbers, and spaces are allowed"),
			validation.RuneLength(3, 200).Error("title must be between 3 and 200 characters"),
		),
		validation.Field(&r.Content,
			validation.Required.Error("content is required"),
			validation.RuneLength(10, 5000).Error("body must be between 10 and 5000 characters"),
		),
		validation.Field(&r.TagIDs,
			validation.Required.Error("at least one tag is required"),
			validation.Each(validation.Required),
		),
		validation.Field(&r.CategoryIDs,
			validation.Required.Error("at least one category is required"),
			validation.Each(validation.Required),
		),
	)
}

func (r UpdatePostRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Title,
			validation.Required.Error("title is required"),
			validation.Match(regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)).Error("only letters, numbers, and spaces are allowed"),
			validation.RuneLength(3, 200).Error("title must be between 3 and 200 characters"),
		),
		validation.Field(&r.Content,
			validation.Required.Error("content is required"),
			validation.RuneLength(10, 5000).Error("body must be between 10 and 5000 characters"),
		),
		validation.Field(&r.TagIDs,
			validation.Required.Error("at least one tag is required"),
			validation.Each(validation.Required),
		),
		validation.Field(&r.CategoryIDs,
			validation.Required.Error("at least one category is required"),
			validation.Each(validation.Required),
		),
	)
}
