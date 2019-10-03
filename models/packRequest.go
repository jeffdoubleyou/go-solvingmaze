package models

type PackRequest struct {
	Items               []*Item   `json:"items,models.Item,omitempty"`
	ShowRates           bool      `json:"showRates,bool,omitempty"`
	EnableCatalog       bool      `json:"enableCatalog,bool,omitempty"`
	ValidateDestination string    `json:"validateDestination,omitempty"`
	Destination         *Address  `json:"destination,models.Address,omitempty"`
	ShipDate            *Delivery `json:"shipDate,models.Delivery,omitempty"`
	VoidFill            int       `json:"voidFill,int,omitempty"`
	Services            []string  `json:"services,[]string,omitempty"`
}

type PackResponse struct {
	Status           *Status
	DestinationMatch *Address
	Services         []*Service
}
