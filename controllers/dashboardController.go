package controllers

import (
	"fmt"
	"myapp/models"

	_ "myapp/docs" // This imports the generated swagger documentation

	"github.com/gofiber/fiber/v2"
	// "github.com/swaggo/fiber-swagger"
	"gorm.io/gorm"
)

// @Summary  GetTblqueMonthlyReport
// @Description display username, number, department, counter, called date, issued time, called time filter by start and end date and department id
// @Tags []models.Tblquelog
// @Success 201
// @Router /dashboard/GetTblqueMonthlyReport [Get]
// 20250220 - display username, number, department, counter, called date, issued time, called time filter by start and end date and department id
func GetTblqueMonthlyReport(c *fiber.Ctx) error {
	var items models.Dashboard
	var itemlog []models.Tblquelog

	if err := c.BodyParser(&items); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	startDate := items.BeginTimestamp
	endDate := items.EndTimestamp
	deptid := items.Deptid

	// currentDate := time.Now().Truncate(24 * time.Hour) // Truncate to midnight of today
	// currentDate := time.Now().Format("2006-01-02")

	db := c.Locals("db").(*gorm.DB)

	if err := db.Where("CONVERT(DATE, created_at) >= ? AND CONVERT(DATE, created_at) <= ? AND session_dept_id =  ?", startDate, endDate, deptid).Find(&itemlog).Error; err != nil {
		return c.Status(500).SendString("Error fetching quelogs")
	}

	return c.JSON(itemlog)
}

// 20250220 - display department, number, called, user, counter filter by selected date
func GetTblqueQueList(c *fiber.Ctx) error {
	var items models.Dashboard
	var itemlog []models.Tblquelog

	if err := c.BodyParser(&items); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	reportDate := items.BeginTimestamp

	// currentDate := time.Now().Truncate(24 * time.Hour) // Truncate to midnight of today
	// currentDate := time.Now().Format("2006-01-02")

	db := c.Locals("db").(*gorm.DB)

	if err := db.Where("CONVERT(DATE, created_at) = ?", reportDate).Find(&itemlog).Error; err != nil {
		return c.Status(500).SendString("Error fetching quelogs")
	}

	return c.JSON(itemlog)
}

// 20250220 - display department,number and counter; filter by username and Date selected
// func GetTblquebyUserTimeStamp(c *fiber.Ctx) error {
func GetTblqueUserReport(c *fiber.Ctx) error {
	var items models.Dashboard
	var itemlog []models.Tblquelog

	if err := c.BodyParser(&items); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	sessioncalleds := items.Called
	username := items.Username

	// currentDate := time.Now().Truncate(24 * time.Hour) // Truncate to midnight of today
	// currentDate := time.Now().Format("2006-01-02")

	db := c.Locals("db").(*gorm.DB)

	// if err := db.Where("CALLED = ? AND CONVERT(DATE, created_at) = ?", called, currentDate).Find(&items).Error; err != nil {
	if err := db.Where("username = ? AND  called IN ?", username, sessioncalleds).Find(&itemlog).Error; err != nil {
		return c.Status(500).SendString("Error fetching quelogs")
	}

	return c.JSON(itemlog)
}

// GetItems retrieves all items
func GetTblquebyTimeStamp(c *fiber.Ctx) error {
	var items models.Dashboard
	var itemlog []models.Tblquelog

	if err := c.BodyParser(&items); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	fmt.Println(items.Deptid)
	fmt.Println(items.Counterid)
	fmt.Println(items.Called)

	sessiondeptids := items.Deptid
	sessioncntrids := items.Counterid
	sessioncalleds := items.Called

	// currentDate := time.Now().Truncate(24 * time.Hour) // Truncate to midnight of today
	// currentDate := time.Now().Format("2006-01-02")

	db := c.Locals("db").(*gorm.DB)

	// if err := db.Where("CALLED = ? AND CONVERT(DATE, created_at) = ?", called, currentDate).Find(&items).Error; err != nil {
	if err := db.Where("session_dept_id IN ? AND session_counter_id IN ? AND called IN ?", sessiondeptids, sessioncntrids, sessioncalleds).Find(&itemlog).Error; err != nil {
		return c.Status(500).SendString("Error fetching quelogs")
	}

	return c.JSON(itemlog)
}
