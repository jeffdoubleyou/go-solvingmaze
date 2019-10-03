package models

import "github.com/jeffdoubleyou/go-solvingmaze/models"

type Package struct {
	Container     string
	Weight        float64
	WeightUnit    string
	Length        float64
	Width         float64
	Height        float64
	DimensionUnit string
	Value         float64
	Items         []*models.Item
	Bundle        int
	Tags          []string
}
