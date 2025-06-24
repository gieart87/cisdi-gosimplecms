package models

import (
	"github.com/gosimple/slug"
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name       string `json:"name"`
	Slug       string `gorm:"size:255;not null;uniqueIndex"`
	Posts      []Post `gorm:"many2many:post_tags"`
	UsageCount int64  `json:"usage_count;default:0"`
}

type TagScore struct {
	TagOne           string  `json:"tag_one"`
	TagTwo           string  `json:"tag_two"`
	Score            float32 `json:"score"`
	CoOccurrentScore int32   `json:"co_occurrent_score"`
}

func (p *Tag) BeforeCreate(tx *gorm.DB) (err error) {
	p.Slug = slug.Make(p.Name)
	return
}
