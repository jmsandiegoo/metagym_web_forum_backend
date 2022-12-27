package dataaccess

import (
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
)

// Saves a new user in database and returns newly created user
func CreateNewUser(user *databasemodels.User) (*databasemodels.User, error) {
	// Generate ID and hash password
	user.UserID = api.GenerateUUID()
	api.PasswordHash(user)

	// store to database
	err := database.Database.Create(&user).Error
	if err != nil {
		return &(databasemodels.User{}), err
	}
	return user, nil
}

func FindUserByUsername(username string) (databasemodels.User, error) {
	var user databasemodels.User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return databasemodels.User{}, err
	}
	return user, nil
}
