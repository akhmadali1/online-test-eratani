// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

import "github.com/shopspring/decimal"

// User -.
type User struct {
	ID             int64  `json:"id"`
	Country        string `json:"country" validate:"required,max=255"`
	CreditCardType string `json:"credit_card_type" validate:"required,max=100"`
	CreditCard     string `json:"credit_card" validate:"required,max=100"`
	FirstName      string `json:"first_name" validate:"required,max=255"`
	LastName       string `json:"last_name" validate:"max=255"`
}

// CountrySpend -.
type CountrySpend struct {
	Country    string          `json:"country"`
	TotalSpend decimal.Decimal `json:"total_spend"`
}

// CreditCardTypeCount -.
type CreditCardTypeCount struct {
	CreditCardType string `json:"credit_card_type"`
	Total          int64  `json:"total"`
}
