// defines the table model and relted operations
// handles restaurant tables and their status

package models

import (
	"errors"

	"gorm.io/gorm"
)

type Table struct {
	gorm.Model
	Number   uint   `gorm:"unique;not null" json:"number"`
	Capacity uint   `gorm:"not null" json:"capacity"`
	Location string `gorm:"size:255" json:"location"`
	Status   string `gorm:"size:50;default:'available'" json:"status"`
}

func (t *Table) BeforeSave(tx *gorm.DB) (err error) {
	if t.Number == 0 || t.Capacity == 0 {
		return errors.New("table and capacity must be grater tha zero")
	}
	return nil
}
