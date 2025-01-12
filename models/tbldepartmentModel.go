package models

import "time"

// Department represents the data structure for the item.
type Tbldepartment struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Letter      string    `json: "letter"`
	StartNumber int       `json: "startnumber"`
	Status      string    `json: status`
	Timestamp   time.Time `json:"timestamp"`
}
