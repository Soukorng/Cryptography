package main
import (
	"fmt"
)
// Bitwise and Assignment Operators

func myXor(a, b int) {
	fmt.Printf("XOR of %d and %d is %d\n", a, b, a ^ b)
}

func myNOT(a, b int) {
	fmt.Printf("not of %d is %d\n", a, ^a)
	fmt.Printf("not of %d is %d\n", b, ^b)
}

func myOR(a, b int) {
	fmt.Printf("OR of %d and %d is %d\n", a, b, a | b)
}

func myAND(a, b int) {
	fmt.Printf("AND of %d and %d is %d\n", a, b, a & b)
}

func leftShift(a, n int) {
	fmt.Printf("Left Shift of %d by %d is %d\n", a, n, a << n)
}

func rightShift(a, n int) {
	fmt.Printf("Right Shift of %d by %d is %d\n", a, n, a >> n)
}

func lab3() {
	fmt.Println("\n===========Lab 3===========")
	fmt.Println("Bitwise and Assignment Operators")
	var a, b, n int

	fmt.Print("\nEnter first number: ")
	fmt.Scanln(&a)
	fmt.Print("Enter second number: ")
	fmt.Scanln(&b)
	fmt.Print("Enter the number of positions to shift: ")
	fmt.Scanln(&n)
	
	myXor(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	leftShift(a, n)
	rightShift(a, n)

}

