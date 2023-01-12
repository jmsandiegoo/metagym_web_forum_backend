package handlers

import (
	"metagym_web_forum_backend/internal/api"
	dataaccess "metagym_web_forum_backend/internal/data-access"
	apimodels "metagym_web_forum_backend/internal/models/api-models"
	databasemodels "metagym_web_forum_backend/internal/models/database-models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
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

	jwt, err := api.GenerateJWT(user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": newUser, "jwt": jwt})
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

	context.JSON(http.StatusOK, gin.H{"user": user, "jwt": jwt})
}

// set ups the profile of a new user
func HandleOnboard(context *gin.Context) {
	var onboardInput apimodels.OnboardInput

	err := context.ShouldBindJSON(&onboardInput)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId, err := api.GetTokenUserId(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currUserProfile, err := dataaccess.FindUserProfileByUserId(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var interests []databasemodels.Interest

	interests, err = dataaccess.FindInterestByIds(onboardInput.Interests)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	profile := databasemodels.UserProfile{
		PfpUrl:     onboardInput.PfpUrl, // TODO File Upload future feature
		Bio:        onboardInput.Bio,
		Experience: databasemodels.Experience_enum(onboardInput.Experience),
		Country:    onboardInput.Country,
		Height:     onboardInput.Height,
		Weight:     onboardInput.Weight,
		Age:        onboardInput.Age,
		Interests:  interests,
	}

	profile.UserID = userId

	if (cmp.Equal(currUserProfile, databasemodels.UserProfile{})) {
		newUserProfile, err := dataaccess.CreateNewUserProfile(&profile)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"profile": newUserProfile})
	} else {
		profile.ID = currUserProfile.ID
		updatedUserProfile, err := dataaccess.UpdateUserProfile(&profile)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		context.JSON(http.StatusOK, gin.H{"profile": updatedUserProfile})
	}
}

// Returns authenticated user data with jwt in request
func HandleGetAuthUser(context *gin.Context) {
	userId, err := api.GetTokenUserId(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currUser, err := dataaccess.FindUserById(userId)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"user": currUser})
}
