package models

import "time"

// Item represents the data structure for the item.
type Sysnotification struct {
	ID               uint      `json:"id"`
	NotificationText string    `json:"notificationtext"`
	Font             string    `json:"font"`
	Color            string    `json:"color"`
	Timestamp        time.Time `json:"timestamp"`
}
