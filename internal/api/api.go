package api

import (
	"encoding/json"

	"github.com/google/uuid"
)

// Types
type Payload struct {
	Meta json.RawMessage `json:"meta,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

type Response struct {
	Payload   Payload  `json:"payload"`
	Messages  []string `json:"messages"`
	ErrorCode int      `json:"errorCode"`
}

// interface for api request input validation
type Validator interface {
	Ok() error
}

// Methods
func GenerateUUID() uuid.UUID {
	return uuid.New()
}
