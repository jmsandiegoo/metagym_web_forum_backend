package dataaccess

import (
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"

	"github.com/google/uuid"
)

// Saves a new user in database and returns newly created user
func CreateNewUser(user *databasemodels.User) (*databasemodels.User, error) {
	// Generate ID and hash password
	user.ID = api.GenerateUUID()
	api.PasswordHash(user)

	// store to database
	err := database.Database.Create(&user).Error
	if err != nil {
		return &(databasemodels.User{}), err
	}
	return user, nil
}

// query user by username
func FindUserByUsername(username string) (databasemodels.User, error) {
	var user databasemodels.User
	err := database.Database.Preload("Profile").Where("username=?", username).Find(&user).Error
	if err != nil {
		return databasemodels.User{}, err
	}
	return user, nil
}

// query user by primary key user_id
func FindUserById(id uuid.UUID) (databasemodels.User, error) {
	var user databasemodels.User
	err := database.Database.Preload("Profile").First(&user, id).Error
	if err != nil {
		return databasemodels.User{}, err
	}
	return user, nil
}

// User profile
func CreateNewUserProfile(userProfile *databasemodels.UserProfile) (*databasemodels.UserProfile, error) {
	userProfile.ID = api.GenerateUUID()
	err := database.Database.Create(&userProfile).Error

	if err != nil {
		return &(databasemodels.UserProfile{}), err
	}

	return userProfile, nil
}

func UpdateUserProfile(userProfile *databasemodels.UserProfile) (*databasemodels.UserProfile, error) {
	err := database.Database.Save(&userProfile).Error

	if err != nil {
		return &(databasemodels.UserProfile{}), err
	}

	return userProfile, nil
}

func FindUserProfileByUserId(id uuid.UUID) (databasemodels.UserProfile, error) {
	var userProfile databasemodels.UserProfile
	err := database.Database.Where("user_id=?", id).Find(&userProfile).Error
	if err != nil {
		return databasemodels.UserProfile{}, err
	}
	return userProfile, nil
}
