package models

import (
	"gorm.io/gorm"
)

type MpesaTransaction struct {
	gorm.Model
	TransactionType string  `json:"transaction_type"` // C2B or B2C
	Amount          float64 `json:"amount"`
	PhoneNumber     string  `json:"phone_number"`
	Status          string  `json:"status"`
	ReceiptNumber   string  `json:"receipt_number"`
	CheckoutID      string
	Description     string
}
