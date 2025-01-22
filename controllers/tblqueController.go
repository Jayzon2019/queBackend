package controllers

import (
	"fmt"
	"myapp/models"
	"reflect"
	"strconv"
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
	var itemlog models.Tblquelog
	var dept models.Tbldepartment
	item.Timestamp = time.Now()

	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	itemlog.QueID = int(item.ID)
	itemlog.Fullname = item.Fullname
	itemlog.Status = item.Status
	itemlog.Called = item.Called
	itemlog.SessionDeptID = item.SessionDeptID
	deptID := int(item.SessionDeptID)

	//increment
	if err := c.Locals("db").(*gorm.DB).First(&dept, deptID).Error; err != nil {
		return c.Status(404).SendString("Item not found")
	}

	if err := c.BodyParser(&dept); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	fmt.Println("item ssdeptid: ", item.SessionDeptID)
	fmt.Println("dept startnumber: ", dept.StartNumber)
	//Convert string to integer
	//num, err := strconv.Atoi(dept.StartNumber)

	num, err := strconv.ParseInt(string(dept.StartNumber), 10, 32)
	if err != nil {
		fmt.Println("Error:", err.Error())
		return c.Status(500).SendString("Error updating start number")
	}

	fmt.Println("dept startnumber 2: ", dept.StartNumber)

	fmt.Println("Type of a:", reflect.TypeOf(dept.StartNumber)) // int
	fmt.Println("Type of a:", reflect.TypeOf(num))              // int

	num = num + 1
	str := fmt.Sprintf("%d", num)

	dept.StartNumber = str
	dept.Timestamp = time.Now()

	fmt.Println("dept startnumber 3: ", dept.StartNumber)
	if err := c.Locals("db").(*gorm.DB).Save(&dept).Error; err != nil {
		return c.Status(500).SendString("Error updating department")
	}
	if err := c.BodyParser(&dept); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}

	//

	//add que
	item.DeptName = dept.Name
	item.Number = dept.StartNumber
	if err := c.Locals("db").(*gorm.DB).Create(&item).Error; err != nil {
		return c.Status(500).SendString("Error creating item")
	}

	//que log
	itemlog.Number = dept.StartNumber
	itemlog.DeptName = dept.Name

	if err := c.Locals("db").(*gorm.DB).Create(&itemlog).Error; err != nil {
		return c.Status(500).SendString("Error creating item")
	}
	return c.Status(201).JSON(item)
}

// UpdateItem updates an existing item by ID
func UpdateTblque(c *fiber.Ctx) error {
	id := c.Params("id")
	var item models.Tblque
	var itemlog models.Tblquelog
	var Tblcounter models.Tblcounter

	item.Timestamp = time.Now()
	if err := c.Locals("db").(*gorm.DB).First(&item, id).Error; err != nil {
		return c.Status(404).SendString("Item not found")
	}
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).SendString("Invalid request body")
	}
	//get countername
	if err := c.Locals("db").(*gorm.DB).First(&Tblcounter, item.SessionCounterID).Error; err != nil {
		return c.Status(404).SendString("Item not found")
	}

	item.CounterName = Tblcounter.Name
	itemlog.CounterName = Tblcounter.Name

	//save item
	if err := c.Locals("db").(*gorm.DB).Save(&item).Error; err != nil {
		return c.Status(500).SendString("Error updating que")
	}

	itemlog.QueID = int(item.ID)
	itemlog.Fullname = item.Fullname
	itemlog.Status = item.Status
	itemlog.Called = item.Called
	itemlog.SessionDeptID = item.SessionDeptID
	itemlog.SessionCounterID = item.SessionCounterID
	itemlog.Number = item.Number
	itemlog.Username = item.Username
	itemlog.Timestamp = item.Timestamp

	if err := c.Locals("db").(*gorm.DB).Create(&itemlog).Error; err != nil {
		return c.Status(500).SendString("Error updating quelog")
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
