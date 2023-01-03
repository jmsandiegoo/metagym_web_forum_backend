package handlers

import (
	"metagym_web_forum_backend/internal/api"
	dataaccess "metagym_web_forum_backend/internal/data-access"
	apimodels "metagym_web_forum_backend/internal/models/api-models"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var signupInput apimodels.SignupInput

	err := context.ShouldBindJSON(&signupInput)

	if err != nil {
		// return error
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := databasemodels.User{
		Username:  signupInput.Username,
		Email:     signupInput.Email,
		FirstName: signupInput.FirstName,
		LastName:  signupInput.LastName,
		Password:  signupInput.Password,
	}

	newUser, err := dataaccess.CreateNewUser(&user)

	if err != nil {
		// return error
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": newUser})
}

func Login(context *gin.Context) {
	var loginInput apimodels.LoginInput

	err := context.ShouldBindJSON(&loginInput)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	user, err := dataaccess.FindUserByUsername(loginInput.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	err = api.ValidatePassword(loginInput.Password, &user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	jwt, err := api.GenerateJWT(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

func Onboard(context *gin.Context) {
	var onboardInput apimodels.OnboardInput

	err := context.ShouldBindJSON(&onboardInput)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := api.GetTokenUserID(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := databasemodels.UserProfile{
		PfpUrl:     onboardInput.PfpUrl, // TODO File Upload
		Bio:        onboardInput.Bio,
		Experience: databasemodels.Experience_enum(onboardInput.Experience),
		Country:    onboardInput.Country,
		Height:     onboardInput.Height,
		Weight:     onboardInput.Weight,
	}

	profile.UserID = userId

	userProfile, err := dataaccess.CreateNewUserProfile(&profile)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"profile": userProfile})

}

// Login

// Reset Password

// Onboard
