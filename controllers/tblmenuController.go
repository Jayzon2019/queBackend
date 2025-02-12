package controllers

import (
	"myapp/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetItems retrieves all items
func GetTblmenu(c *fiber.Ctx) error {
	role := c.Params("id")
	var items []models.Tblmenus

	db := c.Locals("db").(*gorm.DB)

	if err := db.Where("role LIKE ?", role).Order("role, OrderList ").Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching userlogs")
	}

	// if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
	// 	return c.Status(500).SendString("Error fetching items")
	// }
	return c.JSON(items)
}

// GetItems retrieves all items
func GetTblmenus(c *fiber.Ctx) error {
	var items []models.Tblmenus

	db := c.Locals("db").(*gorm.DB)

	if err := db.Order("role, OrderList ").Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching userlogs")
	}

	// if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
	// 	return c.Status(500).SendString("Error fetching items")
	// }
	return c.JSON(items)
}

// CreateItem creates a new item
func CreateMenu(c *fiber.Ctx) error {
	var item models.Tblmenus

	//item.Price = num
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
		//return c.Status(400).JSON(item)
	}
	if err := c.Locals("db").(*gorm.DB).Create(&item).Error; err != nil {
		return c.Status(500).SendString("Error creating item")
	}
	return c.Status(201).JSON(item)
}

// UpdateItem updates an existing item by ID
func UpdateTblmenu(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Tblmenus
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
func DeleteTblmenu(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := c.Locals("db").(*gorm.DB).Delete(&models.Tblmenus{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting item")
	}
	//return c.SendString("Item deleted successfully")
	return c.JSON(id)
}
