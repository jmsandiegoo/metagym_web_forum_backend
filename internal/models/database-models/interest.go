package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type Interest struct {
	InterestID uuid.UUID `gorm:"primaryKey"`
	Name       string
	CreatedAt  time.Time
	UpdatedAt  time.Time

	// Many to many associations
	UserProfiles []UserProfile `gorm:"many2many:user_interests;"`
}
