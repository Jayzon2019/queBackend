package controllers

import (
	"fmt"
	"myapp/models"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

func GenUserReportExcelFile(c *fiber.Ctx) error {
	// Create a new Excel file
	f := excelize.NewFile()

	index := 0
	// Create a sheet called "Users"
	index, err := f.NewSheet("UserReport")

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error retrieving data: %v", err))
	}
	// Set headers for the Excel file
	f.SetCellValue("UserReport", "A1", "#")
	f.SetCellValue("UserReport", "B1", "Username")
	f.SetCellValue("UserReport", "C1", "Department")
	f.SetCellValue("UserReport", "D1", "Number")
	f.SetCellValue("UserReport", "E1", "Counter")

	var items models.Dashboard
	var itemlog []models.Tblquelog

	if err := c.BodyParser(&items); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	sessioncalleds := items.Called
	username := items.Username

	db := c.Locals("db").(*gorm.DB)

	// if err := db.Where("CALLED = ? AND CONVERT(DATE, created_at) = ?", called, currentDate).Find(&items).Error; err != nil {
	if err := db.Where("username = ? AND  called IN ?", username, sessioncalleds).Find(&itemlog).Error; err != nil {
		return c.Status(500).SendString("Error fetching quelogs")
	}

	// Populate the Excel file with user data
	for i, user := range itemlog {
		row := i + 2 // Start from row 2
		f.SetCellValue("UserReport", fmt.Sprintf("A%d", row), i+1)
		f.SetCellValue("UserReport", fmt.Sprintf("B%d", row), user.Username)
		f.SetCellValue("UserReport", fmt.Sprintf("C%d", row), user.DeptName)
		f.SetCellValue("UserReport", fmt.Sprintf("D%d", row), user.Number)
		f.SetCellValue("UserReport", fmt.Sprintf("E%d", row), user.CounterName)
	}

	// Set the active sheet of the Excel file
	f.SetActiveSheet(index)

	// Save the Excel file to a buffer
	if _, err := f.WriteTo(c); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error generating the Excel file: %v", err))
	}

	// Set content type to Excel MIME type
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", "attachment; filename=userreport.xlsx")

	return nil
}

func GenerateExcelFile(c *fiber.Ctx) error {

	// Create a new Excel file
	f := excelize.NewFile()

	index := 0
	// Create a sheet called "Users"
	index, err := f.NewSheet("Users")

	if err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error retrieving data: %v", err))
	}
	// Set headers for the Excel file
	f.SetCellValue("Users", "A1", "ID")
	f.SetCellValue("Users", "B1", "Name")
	f.SetCellValue("Users", "C1", "Email")

	// Populate the Excel file with user data
	// for i, user := range users {
	// 	row := i + 2 // Start from row 2
	// 	f.SetCellValue("Users", fmt.Sprintf("A%d", row), user.ID)
	// 	f.SetCellValue("Users", fmt.Sprintf("B%d", row), user.Name)
	// 	f.SetCellValue("Users", fmt.Sprintf("C%d", row), user.Email)
	// }

	// Set the active sheet of the Excel file
	f.SetActiveSheet(index)

	// Save the Excel file to a buffer
	if _, err := f.WriteTo(c); err != nil {
		return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error generating the Excel file: %v", err))
	}

	// Set content type to Excel MIME type
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", "attachment; filename=users.xlsx")

	return nil
}
