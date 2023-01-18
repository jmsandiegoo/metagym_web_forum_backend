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
	ID         uuid.UUID       `gorm:"primaryKey" json:"userProfileId"`
	Rep        int             `gorm:"not null; default:0" json:"rep"`
	PfpUrl     string          `gorm:"not null" json:"pfpUrl"`
	Bio        string          `gorm:"not null" json:"bio"`
	Experience Experience_enum `gorm:"not null; default:'beginner'" json:"experience"`
	Country    string          `gorm:"not null" json:"country"`
	Height     float32         `gorm:"not null" json:"height"`
	Weight     float32         `gorm:"not null" json:"weight"`
	Age        int             `gorm:"not null" json:"age"`
	CreatedAt  time.Time       `json:"createdAt"`
	UpdatedAt  time.Time       `json:"updatedAt"`

	// Foreign keys
	UserID uuid.UUID `json:"userId"`

	// many to many relationship
	Interests []Interest `gorm:"many2many:user_interests;" json:"interests"`
}
