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
func GetTblquebyTran(c *fiber.Ctx) error {
	var items []models.Tblque

	// Get current date (without the time)
	// currentDate := time.Now().Truncate(24 * time.Hour) // Truncate to midnight of today
	currentDate := time.Now().Format("2006-01-02")

	db := c.Locals("db").(*gorm.DB)
	called := "ENGAGED"

	//if err := db.Where("CONVERT(DATE, created_at) = ?", currentDate).Find(&items).Error; err != nil {
	// Correct query format
	//currentDate := time.Now().Format("2006-01-02")
	if err := db.Where("CALLED = ? AND CONVERT(DATE, created_at) = ?", called, currentDate).Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching userlogs")
	}
	return c.JSON(items)
}

// GetItems retrieves all items
func GetTblque(c *fiber.Ctx) error {
	var items []models.Tblque

	// Get current date (without the time)
	// currentDate := time.Now().Truncate(24 * time.Hour) // Truncate to midnight of today
	currentDate := time.Now().Format("2006-01-02")
	// currentDate := time.Now().Truncate(24 * time.Hour) // Truncate to midnight of today

	db := c.Locals("db").(*gorm.DB)
	// called := "WAITING"

	fmt.Println("CurrentDate:", currentDate)
	// if err := db.Raw(`
	// 		SELECT *
	// 			FROM TBLques
	// 			ORDER BY id DESC
	// 	`).Scan(&names).Error; err != nil {
	//     fmt.Println("Error querying distinct names:", err)
	// } else {
	//     fmt.Println("Distinct names ordered by age:", names)
	// }

	// Query to filter by current date
	if err := db.Where("CONVERT(DATE, created_at) = ?", currentDate).Find(&items).Error; err != nil {
		return c.Status(500).SendString("Error fetching items")
	}

	// if err := c.Locals("db").(*gorm.DB).Find(&items).Error; err != nil {
	// 	return c.Status(500).SendString("Error fetching items")
	// }

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

	fmt.Println("Fullname: ", item.Fullname)
	fmt.Println("ID: ", int(item.SessionDeptID))
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

	fmt.Println("item SessionDeptID: ", item.SessionDeptID)
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
	item.Number = fmt.Sprintf("%04s", dept.StartNumber) //dept.StartNumber
	item.Number = dept.Letter + item.Number

	if err := c.Locals("db").(*gorm.DB).Create(&item).Error; err != nil {
		return c.Status(500).SendString("Error creating item")
	}

	//que log
	//itemlog.Number = dept.StartNumber
	itemlog.Number = item.Number
	itemlog.DeptName = dept.Name
	itemlog.QueID = int(item.ID)

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

	itemlog.DeptName = item.DeptName

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
