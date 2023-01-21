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

func HandleGetThreadComments(context *gin.Context) {
	threadIdStr := context.Param("threadId")
	threadId, err := uuid.Parse(threadIdStr)

	if err != nil {
		context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
		return
	}

	comments, err := dataaccess.FindCommentsByThreadId(threadId)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"comments": comments})
}

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
func HandleUpvoteComment(context *gin.Context) {
	commentIdStr := context.Param("commentId")
	commentId, err := uuid.Parse(commentIdStr)

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

	// handle database query in one transaction here

	tx := database.Database.Begin()

	var comment databasemodels.Comment

	comment, err = dataaccess.FindCommentByIdLocked(commentId, tx)

	if err != nil {
		tx.Rollback()
		context.Error(err)
		return
	}

	// handle database query in one transaction here

	var usersLiked []databasemodels.User

	// check if user already upvoted
	usersLiked, err = dataaccess.FindCommentUsersLikedByIdsLocked(&comment, []uuid.UUID{userId}, tx)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		context.Error(err)
		return
	}

	if (len(usersLiked) > 0 && voteInput.Flag) || (len(usersLiked) == 0 && !voteInput.Flag) {
		tx.Rollback()
		context.Error(api.ErrUser{Message: "Invalid Request", Err: err})
		return
	}

	// check if user is in disliked association if yes then x2 upvote val
	addVoteVal := 10
	var usersDisliked []databasemodels.User

	usersDisliked, err = dataaccess.FindCommentUsersDislikedByIdsLocked(&comment, []uuid.UUID{userId}, tx)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		context.Error(err)
		return
	}

	if len(usersDisliked) > 0 {
		addVoteVal *= 2
	}

	if voteInput.Flag {
		err = dataaccess.AddUsersLikedComment(&comment, &user, tx)

		if err != nil {
			tx.Rollback()
			context.Error(err)
			return
		}

		err = dataaccess.DeleteUsersDislikedComment(&comment, &user, tx)

		if err != nil {
			tx.Rollback()
			context.Error(err)
			return
		}

	} else {
		err = dataaccess.DeleteUsersLikedComment(&comment, &user, tx)

		if err != nil {
			tx.Rollback()
			context.Error(err)
			return
		}
	}

	// fetch thread user if it is not the reqeustor
	if comment.UserID != userId {
		user, err = dataaccess.FindUserById(comment.UserID)

		if err != nil {
			context.Error(err)
			return
		}
	}

	if voteInput.Flag {
		err = dataaccess.AddUserProfileRep(&(user.Profile), addVoteVal, tx)
	} else {
		err = dataaccess.SubtractUserProfileRep(&(user.Profile), addVoteVal, tx)
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

// downvote
func HandleDownvoteComment(context *gin.Context) {
	commentIdStr := context.Param("commentId")
	commentId, err := uuid.Parse(commentIdStr)

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

	// handle database query in one transaction here

	tx := database.Database.Begin()

	var comment databasemodels.Comment

	comment, err = dataaccess.FindCommentByIdLocked(commentId, tx)

	if err != nil {
		tx.Rollback()
		context.Error(err)
		return
	}

	var usersDisliked []databasemodels.User

	// check if user already downvoted
	usersDisliked, err = dataaccess.FindCommentUsersDislikedByIdsLocked(&comment, []uuid.UUID{userId}, tx)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		context.Error(err)
		return
	}

	if (len(usersDisliked) > 0 && voteInput.Flag) || (len(usersDisliked) == 0 && !voteInput.Flag) {
		tx.Rollback()
		context.Error(api.ErrUser{Message: "Invalid Request", Err: err})
		return
	}

	// check if user is in liked association if yes then x2 downvotevote val
	subVoteVal := 10
	var usersLiked []databasemodels.User

	usersLiked, err = dataaccess.FindCommentUsersLikedByIdsLocked(&comment, []uuid.UUID{userId}, tx)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tx.Rollback()
		context.Error(err)
		return
	}

	if len(usersLiked) > 0 {
		subVoteVal *= 2
	}

	if voteInput.Flag {
		err = dataaccess.AddUsersDislikedComment(&comment, &user, tx)

		if err != nil {
			tx.Rollback()
			context.Error(err)
			return
		}

		err = dataaccess.DeleteUsersLikedComment(&comment, &user, tx)

		if err != nil {
			tx.Rollback()
			context.Error(err)
			return
		}
	} else {
		err = dataaccess.DeleteUsersDislikedComment(&comment, &user, tx)

		if err != nil {
			tx.Rollback()
			context.Error(err)
			return
		}
	}

	// fetch thread user if it is not the reqeustor
	if comment.UserID != userId {
		user, err = dataaccess.FindUserById(comment.UserID)

		if err != nil {
			context.Error(err)
			return
		}
	}

	if voteInput.Flag {
		err = dataaccess.SubtractUserProfileRep(&(user.Profile), subVoteVal, tx)
	} else {
		err = dataaccess.AddUserProfileRep(&(user.Profile), subVoteVal, tx)
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

	context.JSON(http.StatusOK, gin.H{})
}
