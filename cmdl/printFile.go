package main

import (
	"fmt"
	"os"
)

func main() {
	// Open the file (or create it if it doesn't exist)
	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	// Close the file at the end of the function
	defer file.Close()

	// Use fmt.Fprintf to write formatted text to the file
	_, err = fmt.Fprintf(file, "Hello, this is a test!\n")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("Text has been written to the file.")
}
