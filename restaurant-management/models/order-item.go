// defines OrderItem models and it's related ops
// handles individual items within an order

package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID uint `gorm:"not null" json:"order_id"`
	Order Order `json:"-"`
	FoodID uint `gorm:"not null" json:"food_id"`
	Food Food `json:"food"`
	Quantity int `gorm:"not null;default:1" json:"quantity"`
	UnitPrice float64 `gorm:"not null" json:"unit_price"`
	Notes string `gorm:"size:512" json:"notes"`
	Status string `gorm:"size:50;default:'pending'" json:"status"`
}