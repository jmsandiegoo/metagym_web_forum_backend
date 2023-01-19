package apimodels

import "github.com/google/uuid"

type SearchThreadInput struct {
	Interests []uuid.UUID `form:"interests"`
	Title     string      `form:"title"`
}
