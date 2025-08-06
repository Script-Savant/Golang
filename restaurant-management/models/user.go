// define the user model and related operations
// handle user data structure and database interactions

package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"size:255;not null" json:"name"`
	Email    string `gorm:"size:255;not null;unique" json:"email"`
	Password string `gorm:"size:255;not null" json:"-"`
	Role     string `gorm:"size:50;default:'waiter';" json:"role"`
}

// BeforeSave - gorm hook that hashes the password before saving
/*
1. Check if the password has been modified
2. Generate a hash from the password
3. Replace the plain pass with the hash
*/
func (u *User) BeforeSave(tx *gorm.DB) error {
	if u.Password == "" {
		return nil
	}
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPass)
	return nil
}

// VerifyPassword - compare the provided password with the stored hash
/*
1. use bccrypt to compare the provided pass with the stored hash
2. return error if they do not match
*/
func (u *User) VerifyPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
