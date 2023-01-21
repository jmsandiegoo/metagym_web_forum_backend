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
		protectedRoutes.GET("/search", handlers.HandleSearch)
		// user routes
		userRoutes := protectedRoutes.Group("/users")
		{
			userRoutes.GET("/auth-user", handlers.HandleGetAuthUser)
			userRoutes.POST("/onboard", handlers.HandleOnboard)
		}

		// interest routes
		interestRoutes := protectedRoutes.Group("/interests")
		{
			interestRoutes.GET("/", handlers.HandleGetAllInterest)
		}

		threadRoutes := protectedRoutes.Group("/threads")
		{
			threadRoutes.GET("/:threadId", handlers.HandleGetThread)
			threadRoutes.POST("/create", handlers.HandleCreateThread)
			threadRoutes.POST("/upvote/:threadId", handlers.HandleUpvoteThread)
			threadRoutes.POST("/downvote/:threadId", handlers.HandleDownvoteThread)
			threadRoutes.PUT("/:threadId", handlers.HandleEditThread)
			threadRoutes.DELETE("/:threadId", handlers.HandleDeleteThread)
		}

		commentRoutes := protectedRoutes.Group("/comments")
		{
			commentRoutes.GET("/:threadId", handlers.HandleGetThreadComments)
			commentRoutes.POST("/create", handlers.HandleCreateComment)
			commentRoutes.POST("/upvote/:commentId", handlers.HandleUpvoteComment)
			commentRoutes.POST("/downvote/:commentId", handlers.HandleDownvoteComment)
			commentRoutes.PUT("/:commentId", handlers.HandleEditComment)
			commentRoutes.DELETE("/:commentId", handlers.HandleDeleteComment)
		}
	}
}
