package databasemodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserInterest struct {
	UserProfileID uuid.UUID `gorm:"primaryKey"`
	InterestID    uuid.UUID `gorm:"primaryKey"`
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
