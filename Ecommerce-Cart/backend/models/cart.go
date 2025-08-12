/*
 Purpose: Cart model and line items
 Pseudo-code:
   - Cart struct (ID, UserID, Items []CartItem)
   - CartItem struct (ID, CartID, ProductID, Qty, PriceAtAdd)
   - helpers: GetTotal
*/
package models

import "time"

// Cart represents a user's shopping cart.
type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	UserID    uint       `gorm:"index" json:"user_id"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

// CartItem is an item in a Cart.
type CartItem struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CartID      uint      `gorm:"index" json:"cart_id"`
	ProductID   uint      `json:"product_id"`
	Quantity    int64     `json:"quantity"`
	PriceAtAdd  int64     `json:"price_at_add"`
}

// GetTotal computes total price (in cents) for the cart.
func (c *Cart) GetTotal() int64 {
	// Steps:
	// 1. Iterate over items and sum PriceAtAdd * Quantity
	// 2. Return total in cents
	var total int64
	for _, it := range c.Items {
		total += it.PriceAtAdd * it.Quantity
	}
	return total
}