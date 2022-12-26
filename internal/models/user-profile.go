package models

import (
	"time"

	"github.com/google/uuid"
)

type experience_enum string

const (
	beginner     experience_enum = "beginner"
	intermediate experience_enum = "experience"
	expert       experience_enum = "expert"
)

type UserProfile struct {
	UserProfileID uuid.UUID       `gorm:"primaryKey"`
	rep           uint            `gorm:"not null; default:0"`
	PfpUrl        string          `gorm:"not null"`
	Bio           string          `gorm:"not null"`
	Experience    experience_enum `gorm:"not null; default:'beginner'"`
	Country       string          `gorm:"not null"`
	Height        float32         `gorm:"not null"`
	Weight        float32         `gorm:"not null"`
	Age           int             `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time

	// Foreign keys
	UserID uuid.UUID
}
