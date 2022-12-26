package routes

import (
	"github.com/gin-gonic/gin"
)

// func GetRoutes() func(r *gin.Engine) {
// 	return func(r chi.Router) {
// 		r.Get("/users", func(w http.ResponseWriter, req *http.Request) {
// 			response, _ := users.HandleList(w, req)

//				w.Header().Set("Content-Type", "application/json")
//				json.NewEncoder(w).Encode(response)
//			})
//		}
//	}
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
