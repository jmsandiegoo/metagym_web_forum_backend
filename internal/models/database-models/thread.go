package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"threadId"`
	Title     string    `gorm:"not null" json:"title"`
	Body      string    `gorm:"not null" json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Foreign keys
	UserID uuid.UUID `json:"userId"`
	User   User      `json:"user"`

	// Has many association
	Comments []Comment `json:"comments"`

	// many to many associations
	Interests     []Interest `gorm:"many2many:thread_interests;" json:"interests"`
	UsersLiked    []User     `gorm:"many2many:post_likes;" json:"usersLiked"`
	UsersDisliked []User     `gorm:"many2many:post_dislikes;" json:"usersDisliked"`
}
