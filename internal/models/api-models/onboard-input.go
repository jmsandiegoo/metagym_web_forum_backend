package apimodels

import "github.com/google/uuid"

type OnboardInput struct {
	PfpUrl     string      `json:"pfpUrl"`
	Bio        string      `json:"bio"`
	Experience string      `json:"experience"`
	Country    string      `json:"country"`
	Height     float32     `json:"height,number"`
	Weight     float32     `json:"weight,number"`
	Age        int         `json:"age,number"`
	Interests  []uuid.UUID `json:"interests"`
}
