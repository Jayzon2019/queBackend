package controllers

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
)

// VideoHandler serves the video file to the client
func VideoHandler(c *fiber.Ctx) error {
	// Get the video filename from URL

	videoFile := c.Params("filename")

	// Ensure the file exists
	if _, err := os.Stat(videoFile); err != nil {
		return c.Status(fiber.StatusNotFound).SendString("Video not found")
	}

	// Open the video file
	file, err := os.Open(videoFile)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Could not open video file")
	}
	defer file.Close()

	// Set headers for video streaming
	c.Set("Content-Type", "video/mp4")
	c.Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", filepath.Base(videoFile)))

	// Serve the video file content to the client
	return c.SendStream(file)
}

func VideoHandler1(c *fiber.Ctx) error {
	// Path to the video file
	videoPath := "./sample.mp4" // Adjust this path to your video file location

	// Open the video file
	file, err := os.Open(videoPath)
	if err != nil {
		return c.Status(500).SendString("Error opening video file")
	}
	defer file.Close()

	// Get file information to set the Content-Length and modify headers accordingly
	fileStats, err := file.Stat()
	if err != nil {
		return c.Status(500).SendString("Error getting file stats")
	}

	// Set headers to inform client this is a video
	c.Set("Content-Type", "video/mp4")
	c.Set("Content-Length", fmt.Sprintf("%d", fileStats.Size()))
	c.Set("Accept-Ranges", "bytes")

	// Serve the video content (similar to streaming)
	return c.SendFile(videoPath)
}

func VideoHandler2(c *fiber.Ctx) error {
	// Path to the video file
	videoPath := "./video2.mp4" // Adjust this path to your video file location

	// Open the video file
	file, err := os.Open(videoPath)
	if err != nil {
		return c.Status(500).SendString("Error opening video file")
	}
	defer file.Close()

	// Get file information to set the Content-Length and modify headers accordingly
	fileStats, err := file.Stat()
	if err != nil {
		return c.Status(500).SendString("Error getting file stats")
	}

	// Set headers to inform client this is a video
	c.Set("Content-Type", "video/mp4")
	c.Set("Content-Length", fmt.Sprintf("%d", fileStats.Size()))
	c.Set("Accept-Ranges", "bytes")

	// Serve the video content (similar to streaming)
	return c.SendFile(videoPath)
}
