package controllers

import (
	"strings"
)

func addLeadingZeros(input string, desiredLength int) string {
	// Calculate how many zeros to prepend
	zerosToAdd := desiredLength - len(input)
	if zerosToAdd > 0 {
		// Prepend zeros
		input = strings.Repeat("0", zerosToAdd) + input
	}
	return input
}

// func main() {
// 	original := "12345"
// 	result := addLeadingZeros(original, 10)
// 	fmt.Println(result) // Output: "0000012345"
// }
