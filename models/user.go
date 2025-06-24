package models

import (
	"database/sql"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string       `gorm:"size:100;not null"`
	Email      string       `gorm:"size:100;not null;uniqueIndex"`
	VerifiedAt sql.NullTime `gorm:"index"`
	Password   string       `gorm:"size:255;not null"`
	Role       string       `gorm:"size:20;not null;index;default:'USER'"`
}

const (
	RoleUser   = "USER"
	RoleAuthor = "AUTHOR"
	RoleEditor = "EDITOR"
	RoleAdmin  = "ADMIN"
)

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r RegisterRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Name, validation.Required, validation.RuneLength(3, 100)),
		validation.Field(&r.Email, validation.Required, validation.Length(5, 50), is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 20)),
	)
}

type RegisterResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (r LoginRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Email, validation.Required, validation.Length(5, 50), is.Email),
		validation.Field(&r.Password, validation.Required, validation.Length(6, 20)),
	)
}
