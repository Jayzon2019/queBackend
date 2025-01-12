package controllers

import (
	"myapp/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetItems retrieves all items
func GetSyscompany(c *fiber.Ctx) error {
	var items []models.Syscompany
	if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching items")
	}
	return c.JSON(items)
}

// CreateItem creates a new item
func CreateSyscompany(c *fiber.Ctx) error {
	var item models.Syscompany
	item.Timestamp = time.Now()
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if err := c.Locals("db").(*gorm.DB).Create(&item).Error; err != nil {
		return c.Status(500).SendString("Error creating item")
	}
	return c.Status(201).JSON(item)
}

// UpdateItem updates an existing item by ID
func UpdateSyscompany(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Syscompany
	item.Timestamp = time.Now()
	if err := c.Locals("db").(*gorm.DB).First(&item, id).Error; err != nil {
		return c.Status(404).SendString("Item not found")
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if err := c.Locals("db").(*gorm.DB).Save(&item).Error; err != nil {
		return c.Status(500).SendString("Error updating item")
	}
	return c.JSON(item)
}

// DeleteItem deletes an item by ID
func DeleteSyscompany(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := c.Locals("db").(*gorm.DB).Delete(&models.Syscompany{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting item")
	}
	//return c.SendString("Item deleted successfully")
	return c.JSON(id)
}
