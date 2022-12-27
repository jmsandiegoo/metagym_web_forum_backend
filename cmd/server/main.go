package main

import (
	"log"
	"metagym_web_forum_backend/internal/database"
	"metagym_web_forum_backend/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	r := gin.Default()
	routes.GetRoutes(r)

	r.Run("localhost:8080")
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func loadDatabase() {
	database.ConnectDb()
	database.AutoMigrateDb()
}
