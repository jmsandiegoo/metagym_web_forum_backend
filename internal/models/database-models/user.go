package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID     uuid.UUID `gorm:"primaryKey"`
	Username   string    `gorm:"not null; unique"`
	Email      string    `gorm:"not null; unique"`
	FirstName  string    `gorm:"not null"`
	LastName   string    `gorm:"not null"`
	Password   string    `gorm:"not null" json:"-"`
	IsVerified bool      `gorm:"not null; default:true"` // change when implementing OTP
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// Has one associations
	Profile UserProfile
}
