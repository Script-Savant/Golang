// Purpose: Define User model and related methods
// Pseudo-code:
// 1. Define User struct with GORM tags
// 2. Add BeforeCreate hook for password hashing
// 3. Add password verification method

package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Addresses []Address `gorm:"foreignKey:UserID" json:"addresses,omitempty"`
	Cart      Cart      `gorm:"foreignKey:UserID" json:"cart,omitempty"`
}

// BeforeCreate hook to hash password before saving user
func (u *User) BeforeSave(tx *gorm.DB) error {
	// 1. Generate hashed password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 2. Set hashed password
	u.Password = string(hashedPassword)
	return nil
}

// VerifyPassword checks if the provided password matches the stored hash
func (u *User) VerifyPassword(password string) error {
	// 1. Compare provided password with stored hash
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
