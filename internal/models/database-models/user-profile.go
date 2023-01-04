package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type Experience_enum string

const (
	beginner     Experience_enum = "beginner"
	intermediate Experience_enum = "experience"
	expert       Experience_enum = "expert"
)

type UserProfile struct {
	UserProfileID uuid.UUID       `gorm:"primaryKey"`
	Rep           uint            `gorm:"not null; default:0"`
	PfpUrl        string          `gorm:"not null"`
	Bio           string          `gorm:"not null"`
	Experience    Experience_enum `gorm:"not null; default:'beginner'"`
	Country       string          `gorm:"not null"`
	Height        float32         `gorm:"not null"`
	Weight        float32         `gorm:"not null"`
	Age           int             `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Foreign keys
	UserID uuid.UUID

	// has many relationship
	Interests []Interest `gorm:"many2many:user_interests;"`
}
