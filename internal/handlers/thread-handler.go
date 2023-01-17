package handlers

import (
	"errors"
	"metagym_web_forum_backend/internal/api"
	dataaccess "metagym_web_forum_backend/internal/data-access"
	"metagym_web_forum_backend/internal/database"
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

// Upvote Thread
func HandleUpvoteThread(context *gin.Context) {
	threadIdStr := context.Param("threadId")
	threadId, err := uuid.Parse(threadIdStr)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	var voteInput apimodels.VoteInput

	err = context.ShouldBindJSON(&voteInput)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	userId, err := api.GetTokenUserId(context)

	if err != nil {
		context.Error(err)
		return
	}

	user, err := dataaccess.FindUserById(userId)

	if err != nil {
		context.Error(err)
		return
	}
	var thread databasemodels.Thread

	thread, err = dataaccess.FindThreadById(threadId)

	if err != nil {
		context.Error(err)
		return
	}

	var usersLiked []databasemodels.User

	// check if user already upvoted
	usersLiked, err = dataaccess.FindThreadUsersLikedByIds(&thread, []uuid.UUID{userId})

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		context.Error(err)
		return
	}

	if (len(usersLiked) > 0 && voteInput.Flag == true) || (len(usersLiked) == 0 && voteInput.Flag == false) {
		context.Error(api.ErrUser{Message: "Invalid Request", Err: err})
		return
	}

	// handle database query in one transaction here

	tx := database.Database.Begin()

	if voteInput.Flag {
		err = dataaccess.AddUsersLikedThread(&thread, &user, tx)
	} else {
		err = dataaccess.DeleteUsersLikedThread(&thread, &user, tx)
	}

	if err != nil {
		tx.Rollback()
		context.Error(err)
		return
	}

	if voteInput.Flag {
		err = dataaccess.AddUserProfileRep(&(user.Profile), 10, tx)
	} else {
		err = dataaccess.SubtractUserProfileRep(&(user.Profile), 10, tx)
	}

	if err != nil {
		tx.Rollback()
		context.Error(err)
		return
	}

	err = tx.Commit().Error

	if err != nil {
		context.Error(err)
		return
	}
	// newThread, err := dataaccess.
	context.JSON(http.StatusOK, gin.H{})
}

func HandleDownvoteThread(context *gin.Context) {

}
