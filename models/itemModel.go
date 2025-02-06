package models

import (
	"gorm.io/gorm"
)

// Item represents the data structure for the item.
type Item struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price string `json:"price"`
	//Status string `json: status`
}

// // Counter represents the data structure for the item.
// type Que struct {
// 	ID         uint   `json:"id"`
// 	Department string `json:"department"`
// 	Number     string `json: number`
// 	Counter    string
// 	Status     string `json: status`
// }

// Initialize the database and automatically create the items table
func InitializeDatabase(db *gorm.DB) {
	db.AutoMigrate(&Item{})
	db.AutoMigrate(&Sysrole{})
	db.AutoMigrate(&Sysstatus{})
	db.AutoMigrate(&Syscompany{})
	db.AutoMigrate(&Sysnotification{})
	db.AutoMigrate(&Tbldepartment{})
	db.AutoMigrate(&Tblcounter{})
	db.AutoMigrate(&Tbluser{})
	db.AutoMigrate(&Tbluserlog{})
	db.AutoMigrate(&Tblque{})
	db.AutoMigrate(&Tblquelog{})
	db.AutoMigrate(&Tblimage{})

}
