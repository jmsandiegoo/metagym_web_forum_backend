package apimodels

import "github.com/google/uuid"

type OnboardInput struct {
	PfpUrl     string      `json:"pfpUrl"`
	Bio        string      `json:"bio"`
	Experience string      `json:"experience"`
	Country    string      `json:"country"`
	Height     float32     `json:"height,string"`
	Weight     float32     `json:"weight,string"`
	Age        int         `json:"age,string"`
	Interests  []uuid.UUID `json:"interests"`
}
