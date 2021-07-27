package models

type Service struct {
	Status       *Status
	Carrier      string
	Service      string
	Value        float64
	Currency     string
	Weight       float64
	WeightUnit   string
	Rate         float64
	ListRate     float64
	FreeShipping bool
	Delivery     *Delivery
	Packing      []*Package
}
