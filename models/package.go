package models

type Package struct {
	Container     string  `json:"container"`
	Weight        float64 `json:"weight"`
	Length        float64 `json:"length"`
	Width         float64 `json:"width"`
	Height        float64 `json:"height"`
	WeightUnit    string  `json:"weightUnit"`
	DimensionUnit string  `json:"dimensionUnit"`
	Rate          float64 `json:"rate"`
	Value         float64 `json:"value"`
	Items         []*Item `json:"items"`
}
