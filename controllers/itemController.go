package controllers

import (
	"myapp/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// GetItems retrieves all items
func GetItems(c *fiber.Ctx) error {
	var items []models.Item
	if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching items")
	}
	return c.JSON(items)
}

// CreateItem creates a new item
func CreateItem(c *fiber.Ctx) error {
	var item models.Item
	//fmt.Println(item.Price)
	//c.SendString(string(item.Price))
	//item.Price = (item.Price)
	// var num int
	//num, err := strconv.Atoi(item.Price)
	//if err != nil {
	//fmt.Println("Error:", err)
	//	return c.Status(400).SendString("Invalid request body: Price")
	//}

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
func UpdateItem(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Item
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
func DeleteItem(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := c.Locals("db").(*gorm.DB).Delete(&models.Item{}, id).Error; err != nil {
		return c.Status(500).SendString("Error deleting item")
	}
	//return c.SendString("Item deleted successfully")
	return c.JSON(id)
}
