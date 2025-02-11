package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dtinth/escpos"
	"github.com/tarm/serial"
)

func main() {
	// Open a serial connection to the thermal printer
	c := &serial.Config{
		Name: "/dev/ttyUSB0", // Update to your serial port (e.g., COM1 on Windows)
		Baud: 9600,
	}
	port, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Initialize the ESC/POS printer
	printer := escpos.NewPrinter(port)

	// Initialize the printer
	err = printer.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Set large font (ESC/POS command for large text)
	// In ESC/POS, setting text size to large typically means selecting Font 2.
	// ESC ( 0x1b ) + '!' ( 0x21 ) + value to select text size.
	// 0x10 is a common code for large text.
	// You can modify the size by adjusting this value.

	// Choose a large font (0x10 is the common ESC/POS code for large font)
	err = printer.SetStyles(escpos.TextSize2x)
	if err != nil {
		log.Fatal(err)
	}

	// Print the large text
	_, err = fmt.Fprintf(port, "Hello, Thermal Printer with Big Font!\n")
	if err != nil {
		log.Fatal(err)
	}

	// Print a new line and cut the paper
	err = printer.NewLine(2)
	if err != nil {
		log.Fatal(err)
	}

	// Perform a cut after printing
	err = printer.Cut()
	if err != nil {
		log.Fatal(err)
	}

	// Optionally, add a delay to allow the printer to finish
	time.Sleep(2 * time.Second)
	fmt.Println("Printed with big font size!")
}
