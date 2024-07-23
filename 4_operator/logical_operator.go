package main

import "fmt"

func exampleLogicalOperator() {
	var x bool

	// AND (และ)
	x = (true && false)
	fmt.Println("true && false =", x) // ผลลัพธ์: false

	// OR (หรือ)
	x = (true || false)
	fmt.Println("true || false =", x) // ผลลัพธ์: true

	// NOT (ไม่)
	x = !true
	fmt.Println("!true =", x) // ผลลัพธ์: false
}
