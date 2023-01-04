package api

import (
	"errors"
	"fmt"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

var privateJWTKey = []byte(os.Getenv("JWT_PRIVATE_KEY"))

func GenerateJWT(user databasemodels.User) (string, error) {
	// token time to live
	tokenTTL, _ := strconv.Atoi(os.Getenv("TOKEN_TTL"))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.UserID,                                                  // user id
		"iat":    time.Now().Unix(),                                            // creation time
		"exp":    time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(), // expiry
	})

	return token.SignedString((privateJWTKey))
}

func ValidateToken(context *gin.Context) error {
	token, err := GetToken(context)
	if err != nil {
		return err
	}
	// type assertion
	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("invalid token provided")
}

func GetTokenUserId(context *gin.Context) (uuid.UUID, error) {
	err := ValidateToken(context)
	if err != nil {
		return uuid.Nil, err
	}

	token, _ := GetToken(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId, _ := uuid.Parse(claims["userId"].(string))

	return userId, nil
}

func GetToken(context *gin.Context) (*jwt.Token, error) {
	// extract the token from the req header 'bearer <JWT>'
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	var signedToken string

	if len(splitToken) == 2 {
		signedToken = splitToken[1]
	} else {
		signedToken = ""
	}

	// decode signed token using private key
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return privateJWTKey, nil
	})

	return token, err
}
