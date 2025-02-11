//go get github.com/dtinth/escpos

package cmdl

import (
	"log"
	"time"

	"github.com/dtinth/escpos"
	"github.com/tarm/serial"
)

func main() {
	// Open a serial connection to the printer
	c := &serial.Config{
		Name: "/dev/ttyUSB0", // Update this to your printer's port
		Baud: 9600,
	}
	port, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer port.Close()

	// Initialize ESC/POS printer
	printer := escpos.NewPrinter(port)

	// Print some text
	err = printer.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Print some simple text
	err = printer.Text("Hello, Thermal Printer!\n")
	if err != nil {
		log.Fatal(err)
	}

	// Print a new line and cut the paper
	err = printer.NewLine(2)
	if err != nil {
		log.Fatal(err)
	}
	err = printer.Cut()
	if err != nil {
		log.Fatal(err)
	}

	// Optional: Set a delay to allow the printer to finish printing
	time.Sleep(2 * time.Second)
}
