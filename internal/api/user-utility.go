package api

import (
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// hash password for secure database storage
func PasswordHash(user *databasemodels.User) error {
	// hash with bcrypt's cost (salt)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	return nil
}

func ValidatePassword(password string, user *databasemodels.User) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

var privateJWTKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user databasemodels.User) (string, error) {
	// token time to live
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.UserID,
		"iat": time.Now().Unix(),
		"eat": time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})

	return token.SignedString((privateJWTKey))
}
