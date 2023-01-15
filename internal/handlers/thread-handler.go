package handlers

import (
	"metagym_web_forum_backend/internal/api"
	dataaccess "metagym_web_forum_backend/internal/data-access"
	apimodels "metagym_web_forum_backend/internal/models/api-models"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleCreateThread(context *gin.Context) {

	var threadInput apimodels.ThreadInput

	err := context.ShouldBindJSON(&threadInput)

	if err != nil {
		// return error
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	userId, err := api.GetTokenUserId(context)

	if err != nil {
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

	user, err := dataaccess.FindUserById(userId)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"thread": apimodels.ThreadResponse{
		Thread: *newThread,
		User:   user,
	}})
}

func HandleGetThread(context *gin.Context) {
	threadIdStr := context.Param("threadId")

	threadId, err := uuid.Parse(threadIdStr)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	thread, err := dataaccess.FindThreadById(threadId)

	if err != nil {
		context.Error(err)
		return
	}

	user, err := dataaccess.FindUserById(thread.UserID)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"thread": apimodels.ThreadResponse{
		Thread: thread,
		User:   user,
	}})
}

func HandleEditThread(context *gin.Context) {
	var threadInput apimodels.ThreadInput
	err := context.ShouldBindJSON(&threadInput)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	threadIdStr := context.Param("threadId")
	threadId, err := uuid.Parse(threadIdStr)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	userId, err := api.GetTokenUserId(context)

	if err != nil {
		context.Error(err)
		return
	}

	thread, err := dataaccess.FindThreadById(threadId)

	if err != nil {
		context.Error(err)
		return
	}

	// check for 403 Todo
	if thread.UserID != userId {
		context.Error(api.ErrNotAuthorized{Err: err})
		return
	}

	interests, err := dataaccess.FindInterestByIds(threadInput.Interests)

	if err != nil {
		context.Error(err)
		return
	}

	thread.Title = threadInput.Title
	thread.Body = threadInput.Body
	thread.Interests = interests

	// Do update
	newThread, err := dataaccess.UpdateThread(&thread)

	if err != nil {
		context.Error(err)
		return
	}

	user, err := dataaccess.FindUserById(thread.UserID)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"thread": apimodels.ThreadResponse{
		Thread: *newThread,
		User:   user,
	}})
}

// Todo handleDeleteThread once comment is done
