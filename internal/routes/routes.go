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

	// auth routess
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
			userRoutes.GET("/auth-user", handlers.HandleGetAuthUser)
			userRoutes.POST("/onboard", handlers.HandleOnboard)
		}

		// interest routes
		interestRoutes := protectedRoutes.Group("/interest")
		{
			interestRoutes.GET("/", handlers.HandleGetAllInterest)
		}

		threadRoutes := protectedRoutes.Group("/thread")
		{
			threadRoutes.GET("/:threadId", handlers.HandleGetThread)
			threadRoutes.POST("/create", handlers.HandleCreateThread)
			threadRoutes.POST("/upvote/:threadId", handlers.HandleUpvoteThread)
			threadRoutes.POST("/downvote/:threadId", handlers.HandleDownvoteThread)
			threadRoutes.PUT("/:threadId", handlers.HandleEditThread)
			// threadRoutes.DELETE("/:threadId")
		}

		commentRoutes := protectedRoutes.Group("/comment")
		{
			commentRoutes.POST("/create", handlers.HandleCreateComment)
			commentRoutes.PUT("/:commentId", handlers.HandleEditComment)
			commentRoutes.DELETE("/:commentId", handlers.HandleDeleteComment)
		}
	}
}
