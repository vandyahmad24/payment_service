package models

import "time"

type Transaction struct {
	Id              string
	Amount          int64
	Currency        string
	PaymentMethodId string
	Description     string
	CustomerID      string
	Status          string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type TransactionResponse struct {
	Id            string
	Amount        int64
	Currency      string
	PaymentMethod string
	Description   string
	CustomerID    string
	Status        string
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
