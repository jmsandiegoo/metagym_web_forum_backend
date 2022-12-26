package routes

import (
	"github.com/gin-gonic/gin"
)

func GetRoutes(r *gin.Engine) {
	// auth
	auth := r.Group("/auth")
	{
		auth.POST("/login")
		auth.POST("/signup")
		auth.POST("/password-reset")
	}

	// user
	user := r.Group(("/user"))
	{
		user.POST("/onboard")
	}
}
