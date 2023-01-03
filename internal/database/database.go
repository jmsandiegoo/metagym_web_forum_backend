// Connect to database script
package database

import (
	"fmt"
	models "metagym_web_forum_backend/internal/models/database-models"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// shared db instance accsible by other packages for querying
var Database *gorm.DB

// start a connection to the specified postgresql db from env file
func ConnectDb() {
	var err error
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", host, username, password, databaseName, port)
	Database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}
}

// conducts migrations to the connected database (built-in GORM functionality)
func AutoMigrateDb() {
	Database.AutoMigrate(&(models.User{}), &(models.UserProfile{}))
}
