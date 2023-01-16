package dataaccess

import (
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateNewThread(thread *databasemodels.Thread) (*databasemodels.Thread, error) {
	// Generate ID and hash password
	thread.ID = api.GenerateUUID()

	// store to database
	err := database.Database.Create(&thread).Error
	if err != nil {
		return &(databasemodels.Thread{}), err
	}
	return thread, nil
}

func FindThreadById(id uuid.UUID) (databasemodels.Thread, error) {
	var thread databasemodels.Thread
	err := database.Database.Preload("UsersLiked").Preload("UsersDisliked").Preload("Interests").Where("id=?", id).Find(&thread).Error

	if err != nil {
		return databasemodels.Thread{}, err
	}

	return thread, nil
}

func UpdateThread(thread *databasemodels.Thread) (*databasemodels.Thread, error) {
	err := database.Database.Save(&thread).Error

	if err != nil {
		return &(databasemodels.Thread{}), err
	}

	return thread, nil
}

// Vote related functions requires a gorm transaction variable

func AddUsersLikedThread(thread *databasemodels.Thread, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&thread).Association("UsersLiked").Append(user)

	if err != nil {
		return err
	}

	err = DeleteUsersDislikedThread(thread, user, tx)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUsersLikedThread(thread *databasemodels.Thread, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&thread).Association("UsersLiked").Delete(*user)

	if err != nil {
		return err
	}

	return nil
}

func AddUsersDislikedThread(thread *databasemodels.Thread, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&thread).Association("UsersDisliked").Append(*user)

	if err != nil {
		return err
	}

	err = DeleteUsersLikedThread(thread, user, tx)

	if err != nil {
		return err
	}

	return nil
}

func DeleteUsersDislikedThread(thread *databasemodels.Thread, user *databasemodels.User, tx *gorm.DB) error {
	err := tx.Model(&thread).Association("UsersDisliked").Delete(user)

	if err != nil {
		return err
	}

	return nil
}

func FindThreadUsersLikedByIds(thread *databasemodels.Thread, ids []uuid.UUID) ([]databasemodels.User, error) {
	var users []databasemodels.User
	err := database.Database.Model(&thread).Where(ids).Association("UsersLiked").Find(&users)

	if err != nil {
		return nil, err
	}

	return users, nil
}
