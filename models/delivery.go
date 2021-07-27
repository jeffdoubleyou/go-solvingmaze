package models

type Delivery struct {
	GuaranteedDaysToDelivery int    `json:"guaranteedDaysToDelivery"`
	Date                     string `json:"date"`
	ScheduledDeliveryTime    string `json:"scheduledDeliveryTime,omitempty"`
	Time                     string `json:"time,omitempty"`
}
