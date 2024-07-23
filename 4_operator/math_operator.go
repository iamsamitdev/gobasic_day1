package main

import "fmt"

func exampleMathOperator() {
	var x int

	// การบวก
	x = 5 + 3
	fmt.Println("5 + 3 =", x) // ผลลัพธ์: 8

	// การลบ
	x = 5 - 3
	fmt.Println("5 - 3 =", x) // ผลลัพธ์: 2

	// การคูณ
	x = 5 * 3
	fmt.Println("5 * 3 =", x) // ผลลัพธ์: 15

	// การหาร
	x = 10 / 2
	fmt.Println("10 / 2 =", x) // ผลลัพธ์: 5

	// การหารเอาเศษ (Modulo)
	x = 10 % 3
	fmt.Println("10 % 3 =", x) // ผลลัพธ์: 1

	// เพิ่มค่า 1
	x = 5
	x++
	fmt.Println("5++ =", x) // ผลลัพธ์: 6

	// ลดค่า 1
	x = 5
	x--
	fmt.Println("5-- =", x) // ผลลัพธ์: 4
}
