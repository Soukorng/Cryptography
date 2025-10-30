package main
import (
	"fmt"
)	

// Mini Calculator
func lab4() {
	fmt.Println("\n===========Lab 4===========")
	fmt.Println("Mini Calculator")
	
	var option, a, b int

	for {
		fmt.Println("\n1). Add (+)")
		fmt.Println("2). Subtract (-)")
		fmt.Println("3). Multiply (*)")
		fmt.Println("4). Divide (/)")
		fmt.Println("5). Modulus (%)")
		fmt.Println("6). Exit")
		fmt.Print("Select an operation (1-6): ")
		fmt.Scanln(&option)

		if option == 6 {
			fmt.Println("Exiting the calculator. Goodbye!")
			return
		}

		if option < 1 || option > 6 {
			fmt.Println("Invalid option. Please select a valid operation (1 - 6).")
			continue // restart the loop
		}

		fmt.Print("Enter first number: ")
		fmt.Scanln(&a)
		fmt.Print("Enter second number: ")
		fmt.Scanln(&b)

		switch option {
		case 1: // Addition
			fmt.Printf("Result: %d + %d = %d\n", a, b, a+b)
		case 2: // Subtraction
			fmt.Printf("Result: %d - %d = %d\n", a, b, a-b)

		case 3: // Multiplication
			fmt.Printf("Result: %d * %d = %d\n", a, b, a*b)

		case 4: // Division
			if b != 0 {
				fmt.Printf("Result: %d / %d = %d\n", a, b, a/b)
			} else {
				fmt.Println("Error: Division by zero is undefined.")
			}
		case 5: // Modulus
			if b != 0 {
				fmt.Printf("Result: %d %% %d = %d\n", a, b, a%b)
			} else {
				fmt.Println("Error: Modulus by zero is undefined.")
			}
		}
	}
}
