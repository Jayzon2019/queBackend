package models

import "time"

// Item represents the data structure for the item.
type Sysnotification struct {
	ID               uint      `json:"id"`
	NotificationText string    `gorm:"type:varchar(200);unique;not null;" json:"notificationtext"`
	Font             string    `gorm:"type:varchar(50);null;" json:"font"`
	Color            string    `gorm:"type:varchar(50);null;" json:"color"`
	Timestamp        time.Time `json:"timestamp"`
}

// Name     string `gorm:"type:varchar(100);not null;" json:"name"` // Limit to 100 characters
// Email    string `gorm:"type:varchar(100);unique;not null;" json:"email"`
// Name  string `json:"name" validate:"required,max=100"`
// Email string `json:"email" validate:"required,email"`
