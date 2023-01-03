package api

import (
	databasemodels "metagym_web_forum_backend/internal/models/database-models"

	"golang.org/x/crypto/bcrypt"
)

// hash password for secure database storage
func PasswordHash(user *databasemodels.User) error {
	// hash with bcrypt's cost (salt)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	return nil
}

func ValidatePassword(password string, user *databasemodels.User) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}
