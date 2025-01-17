package controllers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

func VideoHandler1(c *fiber.Ctx) error {
	// Path to the video file
	videoPath := "./video/sample.mp4" // Adjust this path to your video file location

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
	videoPath := "./video/video2.mp4" // Adjust this path to your video file location

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
