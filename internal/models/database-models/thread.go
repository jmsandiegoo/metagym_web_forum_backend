package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"threadId"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	CreatedAt time.Time
	UpdatedAt time.Time

	// Foreign keys
	UserID uuid.UUID `json:"userId"`

	// Has many association
	Comments []Comment `json:"comments"`

	// many to many associations
	Interests     []Interest `gorm:"many2many:thread_interests;" json:"interests"`
	UsersLiked    []User     `gorm:"many2many:post_likes;" json:"usersLiked"`
	UsersDisliked []User     `gorm:"many2many:post_dislikes;" json:"usersDisliked"`
}
