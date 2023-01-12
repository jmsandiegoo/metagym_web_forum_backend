package apimodels

import "github.com/google/uuid"

type ThreadInput struct {
	Title     string      `json:"title"`
	Body      string      `json:"body"`
	Interests []uuid.UUID `json:"interests"`
}
