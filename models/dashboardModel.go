package models

// Counter represents the data structure for the item.
type Dashboard struct {
	BeginTimestamp string   `json:"begintimestamp" validate:"required"`
	EndTimestamp   string   `json:"endtimestamp" validate:"required"`
	Deptid         []int    `json:"deptid"`
	Counterid      []int    `json:"counterid"`
	Called         []string `json:"called"`
	Username       string   `json:"username"`
}
