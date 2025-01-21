package models

import "time"

// Counter represents the data structure for the item.
type Tblcounter struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;" json:"name" validate:"required"`
	Status    string    `gorm:"type:varchar(50);null;" json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
