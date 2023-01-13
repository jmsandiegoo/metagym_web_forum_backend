package dataaccess

import (
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"

	"github.com/google/uuid"
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
	err := database.Database.Preload("Interests").Where("id=?", id).Find(&thread).Error

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
