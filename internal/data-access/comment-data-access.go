package dataaccess

import (
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"

	"github.com/google/uuid"
)

func CreateNewComment(comment *databasemodels.Comment) (*databasemodels.Comment, error) {
	// Generate ID
	comment.ID = api.GenerateUUID()

	// store to database
	err := database.Database.Create(&comment).Error

	if err != nil {
		return &(databasemodels.Comment{}), err
	}

	return comment, nil
}

func FindCommentById(id uuid.UUID) (databasemodels.Comment, error) {
	var comment databasemodels.Comment // garbage collected once no reference
	err := database.Database.Preload("Thread").Where("id=?", id).Find(&comment).Error

	if err != nil {
		return databasemodels.Comment{}, err
	}

	return comment, nil
}

func UpdateComment(comment *databasemodels.Comment) (*databasemodels.Comment, error) {
	err := database.Database.Save(&comment).Error

	if err != nil {
		return &(databasemodels.Comment{}), err
	}

	return comment, nil
}

func DeleteComment(comment *databasemodels.Comment) error {
	err := database.Database.Delete(&comment).Error

	if err != nil {
		return err
	}

	return nil
}
