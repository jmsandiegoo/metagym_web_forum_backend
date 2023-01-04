package dataaccess

import (
	"metagym_web_forum_backend/internal/database"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"

	"github.com/google/uuid"
)

func FindInterestByIds(ids []uuid.UUID) ([]databasemodels.Interest, error) {
	var interests []databasemodels.Interest
	err := database.Database.Where(ids).Find(&interests).Error

	if err != nil {
		return nil, err
	}

	return interests, nil
}
