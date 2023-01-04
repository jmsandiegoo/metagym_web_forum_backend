package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type UserInterest struct {
	UserProfileID uuid.UUID `gorm:"primaryKey"`
	InterestID    uuid.UUID `gorm:"primaryKey"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
