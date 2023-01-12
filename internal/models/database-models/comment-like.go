package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type CommentLike struct {
	CommentID uuid.UUID `gorm:"primaryKey"`
	UserID    uuid.UUID `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
