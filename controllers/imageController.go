package controllers

import (
	"fmt"
	"myapp/models"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UploadFile(c *fiber.Ctx) error {

	// Parse the multipart form (for file uploads)
	// err := c.Request().ParseMultipartForm(10 < 20) // Limit file size to 10MB

	// if err != nil {
	// 	return c.Status(fiber.StatusBadRequest).SendString("Error parsing form")
	// }

	// Get the file from the form
	file, err := c.FormFile("image")

	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("No file uploaded")
	}

	// Create a unique filename
	ext := filepath.Ext(file.Filename)
	filename := fmt.Sprintf("%d%s", time.Now().Unix(), ext)
	savePath := filepath.Join("static", "uploads", filename)

	// Save the file to the filesystem
	err = c.SaveFile(file, savePath)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving file")
	}

	// Save image metadata to the database (optional)
	image := models.Image{
		Filename: filename,
		Filepath: savePath,
	}

	if err := c.Locals("db").(*gorm.DB).Create(&image).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error saving image to database")
	}

	// Return the file URL as JSON response
	return c.JSON(fiber.Map{
		"message": "File uploaded successfully",
		"url":     fmt.Sprintf("/uploads/%s", filename),
	})
}
