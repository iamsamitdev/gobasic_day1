package main

import (
	"fmt"
)

func main() {

	// This is a comment
	/*
		This is a multi-line comment
		that spans multiple lines
		to explain the code
	*/

	fmt.Println("--------------------")

	fmt.Println("Hello, World!")
	fmt.Println("Welcome to Go Programming Language!")
	fmt.Print("This is a sample Go program ")
	fmt.Print("for the beginners.\n")

	fmt.Println("--------------------")

	fmt.Println(2 + 3) // addition
	fmt.Println(2 - 3) // subtraction
	fmt.Println(2 * 3) // multiplication
	fmt.Println(2 / 3) // division
	fmt.Println(2 % 3) // modulus

	fmt.Println("--------------------")

	// logical AND
	fmt.Println(true && false)
	// logical OR
	fmt.Println(true || false)
	// logical NOT
	fmt.Println(!true)

	fmt.Println("--------------------")

	// bitwise AND
	fmt.Println(2 & 3)
	// bitwise OR
	fmt.Println(2 | 3)
	// bitwise XOR
	fmt.Println(2 ^ 3)
	// bitwise NOT
	fmt.Println(^2)

}
