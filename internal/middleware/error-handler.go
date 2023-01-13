package middleware

import (
	"errors"
	"metagym_web_forum_backend/internal/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()

		for _, err := range context.Errors {
			api.ErrorLogger.Println(err.Err)
			if errors.As(err.Err, &(api.ErrNotAuthenticated{})) {
				context.JSON(http.StatusUnauthorized, gin.H{"error": err.Err.Error()})
				return
			} else if errors.As(err.Err, &(api.ErrUser{})) {
				context.JSON(http.StatusBadRequest, gin.H{"error": err.Err.Error()}) // Todo when doin validation
				return
			} else if errors.Is(err.Err, gorm.ErrRecordNotFound) {
				context.JSON(http.StatusNotFound, gin.H{"error": gorm.ErrRecordNotFound.Error()})
				return
			} else {
				context.JSON(http.StatusInternalServerError, "An error occured. Please try again")
			}
		}
	}
}
