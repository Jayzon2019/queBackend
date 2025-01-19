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

type Tbluserlog struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Username      string    `json:"username"`
	Email         string    `json:"email"`
	Role          string    `json:"role"`
	Password      string    `json:"password"`
	Status        string    `json:"status"`
	Timestamp     time.Time `json:"timestamp"`
	DeptID        int       `json:"deptid"`
	SessionDeptID int       `json:"sessiondeptid"`
}

type Tbluserlogsession struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	Role            string    `json:"role"`
	Password        string    `json:"password"`
	Status          string    `json:"status"`
	Timestamp       time.Time `json:"timestamp"`
	DeptID          int       `json:"deptid"`
	DeptName        string    `json:"deptname"`
	SessionDeptID   int       `json:"deptid"`
	SessionDeptName string    `json:"sessiondeptname"`
}
