package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Name  string `json:"name"`
	Slug  string `gorm:"size:255;not null;uniqueIndex"`
	Posts []Post `gorm:"many2many:post_categories"`
}

func (c *Category) BeforeCreate(tx *gorm.DB) (err error) {
	c.Slug = slug.Make(c.Name)
	return nil
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type UpdateCategoryRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateUpdateCategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (r CreateCategoryRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.RuneLength(3, 100)),
	)
}
