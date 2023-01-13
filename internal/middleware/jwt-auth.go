package middleware

import (
	"metagym_web_forum_backend/internal/api"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		err := api.ValidateToken(context)
		if err != nil {
			context.Error(api.ErrNotAuthenticated{Err: err})
			context.Abort()
			return
		}
		context.Next()
	}
}
