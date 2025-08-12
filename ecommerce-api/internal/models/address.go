// Purpose: Define Address model for user shipping addresses
// Pseudo-code:
// 1. Define Address struct with GORM tags
// 2. Add validation methods if needed

package models

import "time"

type Address struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uint      `gorm:"not null" json:"user_id"`
	Street    string    `gorm:"not null" json:"street"`
	City      string    `gorm:"not null" json:"city"`
	State     string    `gorm:"not null" json:"state"`
	ZipCode   string    `gorm:"not null" json:"zip_code"`
	Country   string    `gorm:"not null" json:"country"`
	IsDefault bool      `gorm:"default:false" json:"is_default"`
}