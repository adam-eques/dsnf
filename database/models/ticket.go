package models

import (
	"time"
)

type Ticket struct {
	ID            uint      `json:"id"`
	Salestime     time.Time `json:"sales_time"`
	DateIssued    string    `json:"date_issued"`
	IsFirstClass  bool      `json:"is_first_class"`
	IsSecondClass bool      `json:"is_second_class"`
	SeatReserved  string    `json:"seat_reserved"`
	Price         float64   `json:"price"`
}
