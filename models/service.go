package models

type Service struct {
	Status struct {
		Success  bool     `json:"success"`
		Messages []string `json:"messages"`
	} `json:"status"`
	Services []struct {
		Carrier    string     `json:"carrier"`
		Service    string     `json:"service"`
		Currency   string     `json:"currency"`
		Weight     float64    `json:"weight"`
		WeightUnit string     `json:"weightUnit"`
		Rate       float64    `json:"rate"`
		Value      int        `json:"value"`
		Packing    []*Package `json:"packing"`
		Status     struct {
			Success bool `json:"success"`
		} `json:"status"`
		ListRate float64   `json:"listRate"`
		Delivery *Delivery `json:"delivery,omitempty"`
	} `json:"services"`
}
