package routes

import (
	"metagym_web_forum_backend/internal/handlers"
	"metagym_web_forum_backend/internal/middleware"
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
	authRoutes := r.Group("/auth")
	{
		authRoutes.POST("/signup", handlers.Signup)
		authRoutes.POST("/login", handlers.Login)
		authRoutes.POST("/password-reset")
	}

	// authentication required routes
	protectedRoutes := r.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	{
		// user routes
		userRoutes := protectedRoutes.Group("/user")
		{
			userRoutes.POST("/onboard", handlers.Onboard)
		}
	}
}
