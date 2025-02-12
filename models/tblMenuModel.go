package models

import "time"

// Department represents the data structure for the item.
type Tblmenus struct {
	ID        uint      `json:"id"`
	Role      string    `gorm:"type:varchar(100);not null;" json:"role" validate:"required"`
	Menu      string    `gorm:"type:varchar(100);not null;" json:"menu" validate:"required"`
	Text      string    `gorm:"type:varchar(100);null;" json:"text"`
	Icon      string    `gorm:"type:varchar(50);null;" json:"icon"`
	Link      string    `gorm:"type:varchar(100);null;" json:"link"`
	Orderlist string    `gorm:"type:varchar(100);null;" json:"orderlist"`
	Timestamp time.Time `gorm:"default:CURRENT_TIMESTAMP"`
}
