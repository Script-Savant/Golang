/*
 Purpose: User model and helper methods
 Pseudo-code:
   - define User struct (ID, Name, Email, PasswordHash)
   - provide HashPassword and VerifyPassword helper methods
*/
package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents an application user.
type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         string         `gorm:"size:255" json:"name"`
	Email        string         `gorm:"size:255;uniqueIndex" json:"email"`
	PasswordHash string         `gorm:"size:255" json:"-"`
}

// HashPassword hashes a plain password using bcrypt and sets PasswordHash.
func (u *User) HashPassword(plain string, cost int) error {
	// Steps:
	// 1. Call bcrypt.GenerateFromPassword with provided cost
	// 2. Set u.PasswordHash and return error if any
	b, err := bcrypt.GenerateFromPassword([]byte(plain), cost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(b)
	return nil
}

// VerifyPassword checks a plain password against the stored hash.
func (u *User) VerifyPassword(plain string) bool {
	// Steps:
	// 1. Call bcrypt.CompareHashAndPassword
	// 2. Return true if match, false otherwise
	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(plain)); err != nil {
		return false
	}
	return true
}