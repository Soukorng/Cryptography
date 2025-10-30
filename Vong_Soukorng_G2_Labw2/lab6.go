package main

import (
	"fmt"
)

// XOR Encrypt function
func xorEncrypt(text string, key byte) []byte {
	encrypted := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		encrypted[i] = text[i] ^ key
	}
	return encrypted
}

// XOR Decrypt function
func xorDecrypt(text string, key byte) string {
	decrypted := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		decrypted[i] = text[i] ^ key
	}
	return string(decrypted)
}

func lab6() {
	fmt.Println("\n=========== Lab 6 ===========")
	fmt.Println("XOR Encryption and Decryption")

	var choice int

	for {
		fmt.Println("\n1) Encrypt")
		fmt.Println("2) Decrypt")
		fmt.Println("3) Exit")
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)

		if choice == 3 {
			fmt.Println("Exiting program. Goodbye!")
			break
		}

		switch choice {
		case 1: // Encryption
			var input string
			var key byte
			fmt.Print("Enter text to encrypt: ")
			fmt.Scanln(&input)
			fmt.Print("Enter key: ")
			fmt.Scanf("%c\n", &key)

			encrypted := xorEncrypt(input, key)
			fmt.Println("Encrypted text:", string(encrypted))

		case 2: // Decryption
			var encryptedInput string
			var key byte
			fmt.Print("Enter text to decrypt: ")
			fmt.Scanln(&encryptedInput)
			fmt.Print("Enter key: ")
			fmt.Scanf("%c\n", &key)

			decrypted := xorDecrypt(encryptedInput, key)
			fmt.Println("Decrypted text:", decrypted)

		default:
			fmt.Println("Invalid option! Please choose (1 - 3).")
		}
	}
}
