package models

import "time"

// Item represents the data structure for the item.
type Sysrole struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;" json:"name" validate:"required"`
	Timestamp time.Time `json:"timestamp"`
}
