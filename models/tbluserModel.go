package models

import (
	"time"
)

// Item represents the data structure for the item.
type Tbluser struct {
	ID        uint      `json:"id"`
	Name      string    `gorm:"type:varchar(100);not null;" json:"name" validate:"required"`
	Username  string    `gorm:"type:varchar(100);not null;" json:"username" validate:"required"`
	Email     string    `gorm:"type:varchar(100);null;" json:"email"`
	Role      string    `gorm:"type:varchar(100);null;" json:"role"`
	Password  string    `gorm:"type:varchar(100);null;" json:"password" `
	Status    string    `gorm:"type:varchar(50);null;" json:"status"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"createat"`
	Timestamp time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timestamp"`

	// Name     string `gorm:"type:varchar(100);not null;" json:"name"` // Limit to 100 characters
	// Email    string `gorm:"type:varchar(100);unique;not null;" json:"email"`
	// Name  string `json:"name" validate:"required,max=100"`
	// Email string `json:"email" validate:"required,email"`
}

type Tbluserlog struct {
	ID                 uint      `json:"id"`
	Username           string    `gorm:"type:varchar(100);not null;" json:"username"`
	Status             string    `gorm:"type:varchar(100);not null;" json:"status"`
	SessionDeptID      int       `json:"sessiondeptid"`
	SessionDeptname    string    `json:"sessiondeptname"`
	SessionCounterID   int       `json:"sessioncounterid"`
	SessionCountername string    `json:"sessioncountername"`
	CreatedAt          time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"createat"`
	Timestamp          time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timestamp"`
	IsActive           bool      `gorm:"default:true"`
}
