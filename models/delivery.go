package models

import (
	"time"
)

type Delivery struct {
	Date                     *time.Time `json:"date,omitempty"`
	Time                     *time.Time `json:"time,omitempty"`
	GuaranteedDaysToDelivery *int       `json:"guaranteedDaysToDelivery,omitempty"`
	ScheduledDeliveryTime    *string    `json:"scheduledDeliveryTime,omitempty"`
}
