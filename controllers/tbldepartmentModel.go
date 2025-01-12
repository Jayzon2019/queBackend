package controllers

import (
	"myapp/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetItems retrieves all items
func GetTbldepartment(c *fiber.Ctx) error {
	var items []models.Tbldepartment
	if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching items")
	}
	return c.JSON(items)
}

// CreateItem creates a new item
func CreateTbldepartment(c *fiber.Ctx) error {
	var item models.Tbldepartment
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
func UpdateTbldepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Tbldepartment
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
func DeleteTbldepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := c.Locals("db").(*gorm.DB).Delete(&models.Tbldepartment{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting item")
	}
	//return c.SendString("Item deleted successfully")
	return c.JSON(id)
}

// CreateItem creates a new item
func IncrementTbldepartmentNumber(c *fiber.Ctx) error {

	id := c.Params("id")
	var item models.Tbldepartment

	if err := c.Locals("db").(*gorm.DB).First(&item, id).Error; err != nil {
		return c.Status(404).SendString("Item not found")
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	item.StartNumber += 1
	item.Timestamp = time.Now()

	if err := c.Locals("db").(*gorm.DB).Save(&item).Error; err != nil {
		return c.Status(500).SendString("Error updating item")
	}

	return c.JSON(item)
}
