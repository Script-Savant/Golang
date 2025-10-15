package models

import (
	"time"

	"gorm.io/gorm"
)

type Snippet struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Content   string `gorm:"not null"`
	ExpiresIn int    `gorm:"not null"`
	ExpiresAt time.Time
	UserID    uint `gorm:"not null"`
	User      User `gorm:"foreignKey:UserID"`
}

func (s *Snippet) BeforeCreate(tx *gorm.DB) (err error) {
	s.ExpiresAt = time.Now().AddDate(0, 0, s.ExpiresIn)
	return
}
