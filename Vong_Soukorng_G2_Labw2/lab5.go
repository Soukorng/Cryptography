package main

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
)

// Binary, Hex, and Base64 Encoding
func lab5() {
	fmt.Println("\n===========Lab 5===========")
	fmt.Println("Binary, Hex, and Base64 Encoding")
	var input string

	fmt.Print("\nEnter a string: ")
	fmt.Scanln(&input)

	// Convert to binary
	BinaryResults := ""
	for i := 0; i < len(input); i++ {
		BinaryResults += fmt.Sprintf("%08b ", input[i])
	}

	// Convert to Hexadecimal
	hexResults := hex.EncodeToString([]byte(input))

	// Convert to Base64
	base64Results := base64.StdEncoding.EncodeToString([]byte(input))

	// Print all results
	fmt.Println("Binary:", BinaryResults)
	fmt.Println("Hexadecimal:", hexResults)
	fmt.Println("Base64:", base64Results)
}
