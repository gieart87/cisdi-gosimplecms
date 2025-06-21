package models

import (
	"database/sql"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         string       `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name       string       `gorm:"size:100;not null"`
	Email      string       `gorm:"size:100;not null;uniqueIndex"`
	VerifiedAt sql.NullTime `gorm:"index"`
	Password   string       `gorm:"size:255;not null"`
	Role       string       `gorm:"size:20;not null;index;default:'USER'"`
	CreatedAt  time.Time    `gorm:"index"`
	UpdatedAt  time.Time    `gorm:"index"`
	DeletedAt  gorm.DeletedAt
}

const (
	RoleUser   = "USER"
	RoleAuthor = "AUTHOR"
	RoleEditor = "EDITOR"
	RoleAdmin  = "ADMIN"
)

func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	return nil
}
