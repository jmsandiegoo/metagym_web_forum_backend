package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID        uuid.UUID `gorm:"primaryKey" json:"commentId"`
	Body      string    `gorm:"not null" json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Foreign Keys
	ThreadID uuid.UUID `json:"threadId"`
	UserID   uuid.UUID `json:"userId"`

	// Has one
	Thread Thread `json:"thread"`
	User   User   `json:"user"`

	// Many to many relationship
	UsersLiked    []User `gorm:"many2many:comment_likes;" json:"usersLiked"`
	UsersDisliked []User `gorm:"many2many:comment_dislikes;" json:"usersDisliked"`
}
