package models

import "time"

// Item represents the data structure for the item.
type Syscompany struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;" json:"name" validate:"required"`
	Username  string    `gorm:"type:varchar(100);not null;" json:"username" validate:"required"`
	Address   string    `gorm:"type:varchar(100);null; json:"address"`
	Phone     string    `gorm:"type:varchar(100);null; json:"phone"`
	Location  string    `gorm:"type:varchar(1000);null; json:"location"`
	Timestamp time.Time `json:"timestamp"`
}

// Name     string `gorm:"type:varchar(100);not null;" json:"name"` // Limit to 100 characters
// Email    string `gorm:"type:varchar(100);unique;not null;" json:"email"`
// Name  string `json:"name" validate:"required,max=100"`
// Email string `json:"email" validate:"required,email"`
