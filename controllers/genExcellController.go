package controllers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/xuri/excelize/v2"
)

func GenerateExcelFile(c *fiber.Ctx) error {

	// Create a new Excel file
	f := excelize.NewFile()

	index := 0
	// Create a sheet called "Users"
	// index, err := f.NewSheet("Users")

	// if err != nil {
	// 	return c.Status(http.StatusInternalServerError).SendString(fmt.Sprintf("Error retrieving data: %v", err))
	// }
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
