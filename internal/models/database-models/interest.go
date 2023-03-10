package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type Interest struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"interestId"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Many to many associations
	UserProfiles []UserProfile `gorm:"many2many:user_interests;" json:"userProfiles"`
	Threads      []Thread      `gorm:"many2many:thread_interests;" json:"threads"`
}
