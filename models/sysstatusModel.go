package models

import "time"

// Item represents the data structure for the item.
type Sysstatus struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
