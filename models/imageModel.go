package models

import "time"

type Tblimage struct {
	ID        uint      `gorm:"primaryKey"`
	Filename  string    `gorm:"size:255"`
	Filepath  string    `gorm:"size:255"`
	Status    string    `gorm:"type:varchar(50);null;" json:"status"`
	Username  string    `gorm:"type:varchar(100);not null;" json:"username"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP;" json:"createat"`
	Timestamp time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timestamp"`
}
