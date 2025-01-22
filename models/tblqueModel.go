package models

import "time"

// Tblque represents the data structure for the item.
type Tblque struct {
	ID            uint      `json:"id"`
	Fullname      string    `gorm:"type:varchar(100);not null;" json:"fullname"`
	Status        string    `gorm:"type:varchar(100);not null;" json:"status"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Timestamp     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timestamp"`
	IsActive      bool      `gorm:"default:true" json:"isactive"`
	SessionDeptID int       `json:"sessiondeptid"`
	DeptName      string    `gorm:"type:varchar(100);not null;" json:"deptname" validate:"required"`
	Called        string    `gorm:"type:varchar(50);default:WAITING;" json:"called"`
	//waiting, serving, served, cancel, noshow
	//WAITING,CALLED, ENGAGED, CANCEL, DEFER:BACK TO WAITING, TRANSFER, SERVED:END
	Number           string `gorm:"type:varchar(100);null;" json:"number"`
	SessionCounterID int    `json:"sessioncounterid"`
	CounterName      string `gorm:"type:varchar(100);" json:"countername"`
	Username         string `gorm:"type:varchar(100);null;" json:"username"`
}

// Tblquelog represents the data structure for the item.
type Tblquelog struct {
	ID            uint      `json:"id"`
	QueID         int       `json:"queid"`
	Fullname      string    `gorm:"type:varchar(100);not null;" json:"fullname"`
	Status        string    `gorm:"type:varchar(100);not null;" json:"status"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Timestamp     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"timestamp"`
	IsActive      bool      `gorm:"default:true"`
	SessionDeptID int       `json:"sessiondeptid"`
	DeptName      string    `gorm:"type:varchar(100);not null;" json:"deptname" validate:"required"`
	Called        string    `gorm:"type:varchar(50);null;" json:"called"`
	//waiting, serving, served, cancel, noshow
	//WAITING,CALLED, ENGAGED, CANCEL, DEFER:BACK TO WAITING, TRANSFER, SERVED:END
	Number           string `gorm:"type:varchar(100);null;" json:"number"`
	SessionCounterID int    `json:"sessioncounterid"`
	CounterName      string `gorm:"type:varchar(100);" json:"countername"`
	Username         string `gorm:"type:varchar(100);null;" json:"username"`
}
