// defines the order model and related operations
// handles restaurant orders and their status

package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	TableID    uint    `gorm:"not null" json:"table_id"`
	Table      Table   `json:"table"`
	UserID     uint    `gorm:"not null" json:"user_id"`
	User       User    `json:"user"`
	Status     string  `gorm:"size:50;default:'pending'" json:"status"`
	TotalPrice float64 `gorm:"not null" json:"total_price"`
	Notes      string  `gorm:"size:1024" json:"notes"`
}
