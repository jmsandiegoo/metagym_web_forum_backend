package main

import (
	"log"
	"metagym_web_forum_backend/internal/api"
	"metagym_web_forum_backend/internal/database"
	"metagym_web_forum_backend/internal/middleware"
	"metagym_web_forum_backend/internal/routes"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	api.InitLoggers()
	loadEnv()
	loadDatabase()
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{os.Getenv("FRONTEND_URL")}
	config.AllowHeaders = append(config.AllowHeaders, "Authorization")
	config.AllowCredentials = true
	r.Use(cors.New(config))
	r.Use(middleware.ErrorHandlerMiddleware())
	api.InfoLogger.Println("Starting server...")
	routes.GetRoutes(r)
	r.Run("localhost:8080")
	api.InfoLogger.Println("Server is now listening on localhost:8080")
}

func loadEnv() {
	err := godotenv.Load(`../../.env`)
	// log.Fatal(err)
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.ConnectDb()
	database.AutoMigrateDb()
}
