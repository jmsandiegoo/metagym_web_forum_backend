package handlers

import (
	"errors"
	"metagym_web_forum_backend/internal/api"
	dataaccess "metagym_web_forum_backend/internal/data-access"
	apimodels "metagym_web_forum_backend/internal/models/api-models"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func HandleCreateThread(context *gin.Context) {

	var threadInput apimodels.ThreadInput

	err := context.ShouldBindJSON(&threadInput)

	if err != nil {
		// return error
		context.Error(err)
		return
	}

	userId, err := api.GetTokenUserId(context)

	if err != nil {
		// return error
		context.Error(err)
		return
	}

	var interests []databasemodels.Interest

	interests, err = dataaccess.FindInterestByIds(threadInput.Interests)

	if err != nil {
		context.Error(err)
		return
	}

	thread := databasemodels.Thread{
		Title:     threadInput.Title,
		Body:      threadInput.Body,
		Interests: interests,
		UserID:    userId,
	}

	newThread, err := dataaccess.CreateNewThread(&thread)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"thread": newThread})
}

func HandleGetThread(context *gin.Context) {
	threadIdStr := context.Param("threadId")

	threadId, err := uuid.Parse(threadIdStr)

	if err != nil {
		context.Error(err)
		return
	}

	newThread, err := dataaccess.FindThreadById(threadId)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			context.AbortWithError(http.StatusNotFound, err)
		}
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"thread": newThread})
}
