package models

import "time"

// Item represents the data structure for the item.
type Tblque struct {
	ID         uint      `json:"id"`
	Department string    `json:"department"`
	Number     string    `json:"number"`
	Called     string    `json:"called"` //waiting, serving, served, cancel, noshow
	Counter    string    `json:"counter"`
	IssueTime  time.Time `json:"issuetime"`
	// ServeStartTime time.Time `json:"servestarttime"`
	// ServeEndTime   time.Time `json:"serveendtime"`
}

// Item represents the data structure for the item.
type Tblquelog struct {
	ID             uint      `json:"id"`
	IDDepartment   int       `json:"iddepartment"`
	Department     string    `json:"department"`
	Number         string    `json:"number"`
	Called         string    `json:"called"`
	Counter        string    `json:"counter"`
	IDUsername     int       `jason:"idusername"`
	Username       string    `jason:"username"`
	IDCounter      int       `jason:"idusername"`
	IssueTime      time.Time `json:"issuetime"`
	ServeStartTime time.Time `json:"servestarttime"`
	ServeEndTime   time.Time `json:"serveendtime"`
	ServedTime     time.Time `json:"servedtime"`
	Timestamp      time.Time `json:"timestamp"`
}
