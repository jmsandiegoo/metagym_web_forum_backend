package api

import (
	"github.com/google/uuid"
)

// interface for api request input validation
type Validator interface {
	Ok() error
}

// Methods
func GenerateUUID() uuid.UUID {
	return uuid.New()
}
