// defines the table model and relted operations
// handles restaurant tables and their status

package models

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Number    int    `gorm:"unique;not null" json:"number"`
	Capacity  int    `gorm:"not null" json:"capacity"`
	Location  string `gorm:"size:255" json:"location"`
	Status    string `gorm:"size:50;default:'available'" json:"status"`
	IsSmoking bool   `gorm:"default:false" json:"is_smoking"`
}
