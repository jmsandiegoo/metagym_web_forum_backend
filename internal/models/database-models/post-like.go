package databasemodels

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostLike struct {
	ThreadID  uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	DeletedAt gorm.DeletedAt
}
