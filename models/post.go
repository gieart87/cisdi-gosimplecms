package models

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID                   string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	User                 User
	UserID               string `gorm:"size:36;index"`
	CurrentPostVersionID string `gorm:"size:36;index"`
	PostVersions         []PostVersion
	CreatedAt            time.Time `gorm:"index"`
	UpdatedAt            time.Time `gorm:"index"`
	DeletedAt            gorm.DeletedAt
}
