// Purpose: Define Product model for ecommerce items
// Pseudo-code:
// 1. Define Product struct with GORM tags
// 2. Add validation methods if needed

package models

import "time"

type Product struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Price       float64   `gorm:"not null" json:"price"`
	Stock       int       `gorm:"not null" json:"stock"`
}