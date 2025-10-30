package main
import (
	"fmt"
)
// Logical Operators
func lab2() {

	fmt.Println("\n===========Lab 2===========")
	fmt.Println("Logical Operators")
	
	var a, b int

	fmt.Print("\nEnter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)

	bothPositive := a > 0 && b > 0
	fmt.Println("Both a and b are positive (a > 0 && b > 0): ", bothPositive)

	oneGreater := a > b || b > a
	fmt.Println("Either a is greater than b or b is greater than a (a > b || b > a): ", oneGreater)

	notEqual := a != b
	fmt.Println("a is not equal to b (a != b): ", notEqual)

}
	