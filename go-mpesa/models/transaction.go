package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CheckoutRequestID  string `gorm:"uniqueIndex"`
	TransactionType    string // c2b, b2c
	Amount             float64
	PhoneNumber        string
	AccountReference   string // for c2b bill ref
	ResultCode         int
	ResultDesc         string
	MpesaReceiptNumber string
	TransactionTime    time.Time
}
