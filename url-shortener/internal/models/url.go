package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	OriginalURL string `gorn:"not null"`
	ShortCode   string `gorm:"uniqueIndex;not null"`
	Clicks      uint
}
