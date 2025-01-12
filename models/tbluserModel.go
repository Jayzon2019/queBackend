package models

import "time"

// Item represents the data structure for the item.
type Tbluser struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Password  string    `json:"password"`
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
}
