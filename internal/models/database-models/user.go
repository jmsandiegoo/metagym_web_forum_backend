package databasemodels

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"primaryKey" json:"userId"`
	Username   string    `gorm:"not null; unique" json:"username"`
	Email      string    `gorm:"not null; unique" json:"email"`
	FirstName  string    `gorm:"not null" json:"firstName"`
	LastName   string    `gorm:"not null" json:"lastName"`
	Password   string    `gorm:"not null" json:"-"`
	IsVerified bool      `gorm:"not null; default:true" json:"isVerified"` // change when implementing OTP
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`

	// Has one associations
	Profile UserProfile `json:"profile"`

	// has many associations
	Threads  []Thread  `json:"threads"`
	Comments []Comment `json:"comments"`

	// many to many associations
	LikedThreads     []Thread  `gorm:"many2many:post_likes;" json:"likedThreads"`
	DislikedThreads  []Thread  `gorm:"many2many:post_dislikes;" json:"dislikedThreads"`
	LikedComments    []Comment `gorm:"many2many:comment_likes;" json:"likedComments"`
	DislikedComments []Comment `gorm:"many2many:comment_dislikes;" json:"dislikedComments"`
}
