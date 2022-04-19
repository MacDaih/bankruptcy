package core

import "time"

type Transaction struct {
	Date     time.Time `json:"date"`
	Amount   float64   `json:"amount"`
	Currency string    `json:"currency"`
	To       string    `json:"to,omitempty"`
}

type Group struct {
	Transactions []Transaction `json:"transactions"`
}
