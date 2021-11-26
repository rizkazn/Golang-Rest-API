package models

import (
	"time"
)

type History struct {
	Id         int    `json:"id,omitempty"`
	No_invoice string `json:"no_invoice"`
	Cashier    string `json:"cashier"`
	Date       string `json:"date"`
	Orders     string `json:"orders"`
	Amount     int    `json:"amount"`
}

type Histories []History

func CreateHistory() *History {
	return &History{
		Date: time.Now().Format("2006-02-01"),
	}
}
