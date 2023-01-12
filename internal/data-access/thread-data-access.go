package dataaccess

import (
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
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
