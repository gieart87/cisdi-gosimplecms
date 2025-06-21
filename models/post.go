package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID          string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User        User
	UserID      string       `gorm:"size:36;index"`
	Title       string       `gorm:"size:255"`
	Slug        string       `gorm:"size:255;uniqueIndex"`
	Body        string       `gorm:"type:text"`
	Status      string       `gorm:"size:20;not null;index;default:'DRAFT'"`
	PublishedAt sql.NullTime `gorm:"index"`
	CreatedAt   time.Time    `gorm:"index"`
	UpdatedAt   time.Time    `gorm:"index"`
	DeletedAt   gorm.DeletedAt
}
