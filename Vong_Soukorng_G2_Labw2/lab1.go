package main
import (
	"fmt"
)
// Assignment Operators
func lab1() {

	fmt.Println("\n===========Lab 1===========")
	fmt.Println("Assignment Operators")
	var a, b int

	fmt.Print("\nEnter first number: ")
	fmt.Scanln(&a);

	fmt.Print("Enter second number: ")
	fmt.Scanln(&b);

	c := a
	fmt.Println("assignment Operation (c = a) :", c)

	c += b
	fmt.Println("Addition Assignment Operation (c += b) :", c)

	c -= b
	fmt.Println("Subtraction Assignment Operation (c -= b) :", c)

	c *= b
	fmt.Println("Multiplication Assignment Operation (c *= b) :", c)

	if b != 0 {
		c /= b
		fmt.Println("Division Assignment Operation (c /= b) :", c)
	} else {
		fmt.Println ("Undefined: Division by zero")
	}
	
	if b != 0 {
		c %= b
		fmt.Println("Modulus Assignment Operation (c %= b) :", c)
	} else {
		fmt.Println ("Undefined: Modulus by zero")
	}
}