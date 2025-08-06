// defines an invoice model and related ops
// handles billing and payment information

package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	OrderID       uint       `gorm:"not null" json:"order_id"`
	Order         Order      `json:"order"`
	TotalAmount   float64    `gorm:"not null" json:"total_amount"`
	TaxAmount     float64    `gorm:"not null" json:"tax_amount"`
	Discount      float64    `gorm:"default:0" json:"discount"`
	PaymentMethod string     `gorm:"size:50;default:'cash'" json:"payment_method"`
	IsPaid        bool       `gorm:"default:false" json:"is_paid"`
	PaidAt        *time.Time `json:"paid_at"`
}
