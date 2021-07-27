package models

type Package struct {
	Container     string  `json:"container"`
	Weight        float64 `json:"weight"`
	Length        int     `json:"length"`
	Width         int     `json:"width"`
	Height        int     `json:"height"`
	WeightUnit    string  `json:"weightUnit"`
	DimensionUnit string  `json:"dimensionUnit"`
	Rate          int     `json:"rate"`
	Value         int     `json:"value"`
	Items         []*Item `json:"items"`
}
