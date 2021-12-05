package server

type GPS struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
	Altitude  float64 `json:"alt"`
	Speed     float64 `json:"spd"`
}
