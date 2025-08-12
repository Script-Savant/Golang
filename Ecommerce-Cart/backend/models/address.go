/*
 Purpose: Shipping addresses tied to users
 Pseudo-code:
   - define Address struct (ID, UserID, Line1, Line2, City, State, Zip, Country, IsDefault)
   - belongs-to User
*/
package models

import "time"

// Address represents a user's shipping address.
type Address struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `gorm:"index" json:"user_id"`
	Line1     string    `gorm:"size:255" json:"line1"`
	Line2     string    `gorm:"size:255" json:"line2"`
	City      string    `gorm:"size:100" json:"city"`
	State     string    `gorm:"size:100" json:"state"`
	Zip       string    `gorm:"size:30" json:"zip"`
	Country   string    `gorm:"size:100" json:"country"`
	IsDefault bool      `json:"is_default"`
}