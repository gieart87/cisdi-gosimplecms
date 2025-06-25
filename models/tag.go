package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/gosimple/slug"
	"gorm.io/gorm"
	"regexp"
)

type Tag struct {
	gorm.Model
	Name       string `gorm:"size:255;not null;" json:"name,omitempty"`
	Slug       string `gorm:"size:255;not null;uniqueIndex" json:"slug,omitempty"`
	Posts      []Post `gorm:"many2many:post_tags" json:"-"`
	UsageCount int64  `gorm:"usage_count;default:0"`
}

type TagScore struct {
	TagOne           string  `json:"tag_one"`
	TagTwo           string  `json:"tag_two"`
	Score            float32 `json:"score"`
	CoOccurrentScore int32   `json:"co_occurrent_score"`
}

type TagScorePost struct {
	Tag1ID uint
	Tag2ID uint
	Score  float64
}

func (p *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	p.Slug = slug.Make(p.Name)
	return
}

type CreateTagRequest struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type CreateUpdateTagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func (r CreateTagRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name,
			validation.Required, validation.RuneLength(3, 100),
			validation.Match(regexp.MustCompile(`^[a-zA-Z0-9 ]+$`)).Error("only letters, numbers, and spaces are allowed"),
		),
	)
}
