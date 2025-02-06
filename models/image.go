package models

type Image struct {
	ID       uint   `gorm:"primaryKey"`
	Filename string `gorm:"size:255"`
	Filepath string `gorm:"size:255"`
}
