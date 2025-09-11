package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Content   string `gorm:"type:text;not null"`
	ImagePath string
	AuthorID  uint `gorm:"not null"`
	Author    User `gorm:"foreignKey:AuthorID"`
}
