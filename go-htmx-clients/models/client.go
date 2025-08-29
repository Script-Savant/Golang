package models

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name  string `gorm:"not null" json:"name"`
	Email string `gorm:"unique;not null" json:"email"`
	Phone string `json:"phone"`
}
