package models

import "time"

// Counter represents the data structure for the item.
type Tblcounter struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Status    string    `json: "status"`
	Timestamp time.Time `json:"timestamp"`
}
