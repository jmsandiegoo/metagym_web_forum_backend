package handlers

import (
	"fmt"
	"metagym_web_forum_backend/internal/api"
	dataaccess "metagym_web_forum_backend/internal/data-access"
	apimodels "metagym_web_forum_backend/internal/models/api-models"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCreateThread(context *gin.Context) {

	var threadInput apimodels.ThreadInput

	err := context.ShouldBindJSON(&threadInput)

	if err != nil {
		// return error
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("%+v", threadInput)
	userId, err := api.GetTokenUserId(context)

	if err != nil {
		// return error
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var interests []databasemodels.Interest

	interests, err = dataaccess.FindInterestByIds(threadInput.Interests)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("Intersts gotten: %v", interests)
	// for _, s := range interests {
	// 	fmt.Printf("%+v", s)
	// }

	thread := databasemodels.Thread{
		Title:     threadInput.Title,
		Body:      threadInput.Body,
		Interests: interests,
		UserID:    userId,
	}

	newThread, err := dataaccess.CreateNewThread(&thread)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"thread": newThread})
}
