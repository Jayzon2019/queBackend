package models

import "time"

// Item represents the data structure for the item.
type Syscompany struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	Location  string    `json:"location"`
	Timestamp time.Time `json:"timestamp"`
}
