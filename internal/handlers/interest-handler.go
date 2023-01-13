package handlers

import (
	dataaccess "metagym_web_forum_backend/internal/data-access"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleGetAllInterest(context *gin.Context) {

	interests, err := dataaccess.FindAllInterest()

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"interests": interests})
}
