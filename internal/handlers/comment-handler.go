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

func HandleCreateComment(context *gin.Context) {

	var commentInput apimodels.CommentInput
	err := context.ShouldBindJSON(&commentInput)

	if err != nil {
		// return error
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	threadId, err := uuid.Parse(commentInput.ThreadID)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	var thread databasemodels.Thread

	thread, err = dataaccess.FindThreadById(threadId)

	if err != nil {
		context.Error(err)
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

	// new comment instance
	comment := databasemodels.Comment{
		Body:     commentInput.Body,
		ThreadID: threadId,
		UserID:   userId,
		Thread:   thread,
		User:     user,
	}

	// create new comment
	newComment, err := dataaccess.CreateNewComment(&comment)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"comment": *newComment})
}

func HandleEditComment(context *gin.Context) {
	var commentUpdateInput apimodels.CommentEditInput

	err := context.ShouldBindJSON(&commentUpdateInput)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	commentIdStr := context.Param("commentId")
	commentId, err := uuid.Parse(commentIdStr)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	userId, err := api.GetTokenUserId(context)

	if err != nil {
		context.Error(err)
		return
	}

	comment, err := dataaccess.FindCommentById(commentId)

	if err != nil {
		context.Error(err)
		return
	}

	// check for 403
	if comment.UserID != userId {
		context.Error(api.ErrNotAuthorized{Err: err})
		return
	}

	comment.Body = commentUpdateInput.Body

	// Do update
	newComment, err := dataaccess.UpdateComment(&comment)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"comment": *newComment})
}

func HandleDeleteComment(context *gin.Context) {
	commentIdStr := context.Param("commentId")
	commentId, err := uuid.Parse(commentIdStr)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	userId, err := api.GetTokenUserId(context)

	if err != nil {
		context.Error(err)
		return
	}

	var comment databasemodels.Comment

	comment, err = dataaccess.FindCommentById(commentId)

	if err != nil {
		context.Error(err)
		return
	}

	// check for 403
	if comment.UserID != userId {
		context.Error(api.ErrNotAuthorized{Err: err})
		return
	}

	err = dataaccess.DeleteComment(&comment)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{})
}

// upvote

// downvote
