package models

import "time"

// Item represents the data structure for the item.
type Sysrole struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
}
