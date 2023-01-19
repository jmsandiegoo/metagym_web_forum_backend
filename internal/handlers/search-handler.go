package handlers

import (
	"metagym_web_forum_backend/internal/api"
	dataaccess "metagym_web_forum_backend/internal/data-access"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// support user search in the future TODO
// only search for thread for now
func HandleSearch(context *gin.Context) {
	// query string params
	interestsStr, ok := context.GetQueryArray("interests")
	title := context.Query("title")
	// convert srtring uuid to actual uuid type
	var interests []uuid.UUID
	if ok {
		for _, v := range interestsStr {
			id, err := uuid.Parse(v)

			if err != nil {
				context.Error(api.ErrUser{Message: "Invalid User Request", Err: err})
				return
			}
			interests = append(interests, id)
		}
	}

	threads, err := dataaccess.FindThreadByInterestAndTitle(interests, title)

	if err != nil {
		context.Error(err)
		return
	}

	context.JSON(http.StatusOK, gin.H{"result": threads})
}
