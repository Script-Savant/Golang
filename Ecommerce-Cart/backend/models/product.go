/*
 Purpose: Product model (simple)
 Pseudo-code:
   - Product struct (ID, SKU, Name, PriceCents, Quantity)
*/
package models

import "time"

// Product represents an item that can be added to cart.
type Product struct {
	ID         uint      `gorm:"primaryKey" json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	SKU        string    `gorm:"size:100;uniqueIndex" json:"sku"`
	Name       string    `gorm:"size:255" json:"name"`
	PriceCents int64     `json:"price_cents"`
	Quantity   int64     `json:"quantity"`
}