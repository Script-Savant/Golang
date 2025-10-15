package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username     string    `gorm:"type:varchar(100);not null;uniqueIndex"`
	PasswordHash string    `gorm:"not null"`
	Snippets     []Snippet `gorm:"foreignKey:UserID"`
}
