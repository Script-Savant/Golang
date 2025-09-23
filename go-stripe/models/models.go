package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"not null;uniqueIndex"`
	Password  string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}

type Transaction struct {
	gorm.Model
	UserID          uint   `gorm:"not null"`
	RecipientEmail  string `gorm:"not null"`
	Amount          int64  `gorm:"not null"`
	Currency        string `gorm:"desfault:'usd'"`
	StripePaymentID string `gorm:"uniqueIndex"`
	Status          string `gorm:"deafult:'pending'"`
	User            User   `gorm:"foreignKey:UserID"`
}

func (u *User) HashPassword() error {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPass)
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
