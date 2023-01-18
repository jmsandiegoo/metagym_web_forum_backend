package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type PostDislike struct {
	ThreadID  uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
