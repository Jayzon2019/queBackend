package controllers

import (
	"myapp/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetItems retrieves all items
func GetTbluserlog(c *fiber.Ctx) error {
	var items []models.Tbluserlog
	if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching userlog")
	}
	return c.JSON(items)
}

// CreateItem creates a new item
func CreateTbluserlog(c *fiber.Ctx) error {
	var item models.Tbluserlog
	item.Timestamp = time.Now()
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if err := c.Locals("db").(*gorm.DB).Create(&item).Error; err != nil {
		return c.Status(500).SendString("Error creating userlog")
	}
	return c.Status(201).JSON(item)
}

// UpdateItem updates an existing item by ID
func UpdateTbluserlog(c *fiber.Ctx) error {
	var item models.Tbluserlog

	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	//item2 := item
	item.Timestamp = time.Now()
	username := item.Username

	db := c.Locals("db").(*gorm.DB)

	// Correct query format
	if err := db.Where("username = ?", username).First(&item).Error; err != nil {
		return c.Status(500).SendString("Error fetching userlogs")
	}

	// if err := c.Locals("db").(*gorm.DB).Save(&item2).Error; err != nil {
	// 	return c.Status(500).SendString("Error updating item")
	// }
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if err := db.Save(&item).Error; err != nil {
		return c.Status(500).SendString("Error updating user")
	}

	return c.JSON(item)
}

// GetItems retrieves all items
func GetTbluser(c *fiber.Ctx) error {
	var items []models.Tbluser
	if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching user")
	}
	return c.JSON(items)
}

// CreateItem creates a new item
func CreateTbluser(c *fiber.Ctx) error {
	var item models.Tbluser
	item.Timestamp = time.Now()
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if err := c.Locals("db").(*gorm.DB).Create(&item).Error; err != nil {
		return c.Status(500).SendString("Error creating user")
	}
	return c.Status(201).JSON(item)
}

// UpdateItem updates an existing item by ID
func UpdateTbluser(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Tbluser
	item.Timestamp = time.Now()
	if err := c.Locals("db").(*gorm.DB).First(&item, id).Error; err != nil {
		return c.Status(404).SendString("Item not found")
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if err := c.Locals("db").(*gorm.DB).Save(&item).Error; err != nil {
		return c.Status(500).SendString("Error updating user")
	}

	return c.JSON(item)
}

// DeleteItem deletes an item by ID
func DeleteTbluser(c *fiber.Ctx) error {

	id := c.Params("id")
	if err := c.Locals("db").(*gorm.DB).Delete(&models.Tbluser{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting user")
	}
	//return c.SendString("Item deleted successfully")
	return c.JSON(id)
}

// ValidateUser checks user credentials in the MSSQL database
func ValidateUser(c *fiber.Ctx) error {
	var item models.Tbluser
	var userlog models.Tbluserlog

	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	username := item.Username
	password := item.Password

	db := c.Locals("db").(*gorm.DB)

	// Correct query format
	if err := db.Where("username = ? AND password = ?", username, password).First(&item).Error; err != nil {
		return c.Status(500).SendString("Error fetching users")
	}
	item.Password = "--"
	userlog.Username = item.Username
	userlog.Status = item.Status
	userlog.CreatedAt = item.CreatedAt
	userlog.Timestamp = item.Timestamp
	userlog.IsActive = true

	if err := c.Locals("db").(*gorm.DB).Create(&userlog).Error; err != nil {
		return c.Status(500).SendString("Error creating userlog")
	}

	return c.JSON(item)
}
