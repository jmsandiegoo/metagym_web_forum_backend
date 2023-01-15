package apimodels

import (
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
)

type ThreadResponse struct {
	databasemodels.Thread                     // embedding
	User                  databasemodels.User `json:"user"`
}
