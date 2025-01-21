package models

import "time"

// Department represents the data structure for the item.
type Tbldepartment struct {
	ID          uint      `json:"id"`
	Name        string    `gorm:"type:varchar(100);not null;" json:"name" validate:"required"`
	Letter      string    `gorm:"type:varchar(100);null;" json:"letter"`
	StartNumber string    `gorm:"type:varchar(100);null;" json:"startnumber"`
	Status      string    `gorm:"type:varchar(50);null;" json:"status"`
	Timestamp   time.Time `json:"timestamp"`
}
