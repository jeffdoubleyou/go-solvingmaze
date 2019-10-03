package models

type Prepackage struct {
	Id            int `json:"integer,omitempty"`
	Length        float64
	Width         float64
	Height        float64
	DimensionUnit string  `json:"string,omitempty"`
	Weight        float64 `json:"float64,omitempty"`
	Capacity      int
}
