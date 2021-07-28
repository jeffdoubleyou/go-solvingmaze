package models

type Service struct {
	Carrier    string     `json:"carrier"`
	Service    string     `json:"service"`
	Currency   string     `json:"currency"`
	Weight     float64    `json:"weight"`
	WeightUnit string     `json:"weightUnit"`
	Rate       float64    `json:"rate"`
	Value      float64    `json:"value"`
	Packing    []*Package `json:"packing"`
	Status     struct {
		Success bool `json:"success"`
	} `json:"status"`
	ListRate float64   `json:"listRate"`
	Delivery *Delivery `json:"delivery,omitempty"`
}
