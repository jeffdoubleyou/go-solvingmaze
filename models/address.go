package models

type Address struct {
	FirstName     string   `json:"firstName,omitempty"`
	LastName      string   `json:"lastName,omitempty"`
	Company       string   `json:"company,omitempty"`
	Phone         string   `json:"phone,omitempty"`
	StreetAddress []string `json:"streetAddress,[]models.Address,omitempty"`
	City          string   `json:"city,omitempty"`
	Post          string   `json:"post,omitempty"`
	Region        string   `json:"region,omitempty"`
	Country       string   `json:"country,omitempty"`
	Residential   bool     `json:"residential,bool"`
}
