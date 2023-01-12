package handlers

import (
	dataaccess "metagym_web_forum_backend/internal/data-access"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetAllInterest(context *gin.Context) {

	interests, err := dataaccess.FindAllInterest()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"interests": interests})
}
