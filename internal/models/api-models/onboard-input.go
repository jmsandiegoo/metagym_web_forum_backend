package apimodels

type OnboardInput struct {
	PfpUrl     string  `json:"username"`
	Bio        string  `json:"bio"`
	Experience string  `json:"experience"`
	Country    string  `json:"country"`
	Height     float32 `json:"height"`
	Weight     float32 `json:"weight"`
	Age        string  `json:"age"`
}
