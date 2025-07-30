package models

import "gorm.io/gorm"

// user model
type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Email string `gorm:"unique;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Profile Profile `gorm:"foreignKey:UserID" json:"profile"`
	Post []Post `gorm:"foreignKey:AuthorID" json:"posts"`
}

// profile model for user
type Profile struct {
	gorm.Model
	UserID uint `gorm:"not null" json:"user_id"`
	Bio string `json:"bio"`
	Image string `json:"image"`
	Location string `json:"location"`
}