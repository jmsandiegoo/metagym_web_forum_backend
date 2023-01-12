package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	CommentID uuid.UUID `gorm:"primaryKey"`
	Body      string    `gorm:"not null" json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`

	// Foreign Keys
	ThreadID uuid.UUID `json:"threadId"`
	UserID   uuid.UUID `json:"userId"`

	// Has one
	Thread Thread `gorm:"foreignKey:ThreadID;references:ThreadID" json:"thread"`
	User   Thread `gorm:"foreignKey:UserID;references:UserID" json:"user"`

	// Many to many relationship
	UsersLiked    []User `gorm:"many2many:comment_likes;" json:"usersLiked"`
	UsersDisliked []User `gorm:"many2many:comment_dislikes;" json:"usersDisliked"`
}
