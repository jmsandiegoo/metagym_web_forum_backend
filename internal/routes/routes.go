package routes

import (
	"metagym_web_forum_backend/internal/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Meta Gym Web Forum Api is Running!",
		})
	})

	// auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/signup", handlers.Signup)
		auth.POST("/login", handlers.Login)
		auth.POST("/password-reset")
	}

	// user routes
	user := r.Group(("/user"))
	{
		user.POST("/onboard")
	}
}
