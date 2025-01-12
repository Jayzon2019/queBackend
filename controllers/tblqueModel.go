package controllers

import (
	"myapp/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetItems retrieves all items
func GetTblque(c *fiber.Ctx) error {
	var items []models.Tblque
	if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching items")
	}
	return c.JSON(items)
}

// CreateItem creates a new item
func CreateTblque(c *fiber.Ctx) error {
	var item models.Tblque
	item.IssueTime = time.Now()
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	if err := c.Locals("db").(*gorm.DB).Create(&item).Error; err != nil {
		return c.Status(500).SendString("Error creating item")
	}
	return c.Status(201).JSON(item)
}

// UpdateItem updates an existing item by ID
func UpdateTblque(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Tblque
	item.IssueTime = time.Now()
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
func DeleteTblque(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := c.Locals("db").(*gorm.DB).Delete(&models.Tblque{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting item")
	}
	//return c.SendString("Item deleted successfully")
	return c.JSON(id)
}
