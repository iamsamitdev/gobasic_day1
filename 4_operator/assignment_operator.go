package main

import "fmt"

func exampleAssignmentOperator() {
	var x int

	// กำหนดค่า
	x = 5
	fmt.Println("x = 5 ->", x) // ผลลัพธ์: 5

	// เพิ่มค่าแล้วกำหนดค่า
	x += 3
	fmt.Println("x += 3 ->", x) // ผลลัพธ์: 8

	// ลดค่าแล้วกำหนดค่า
	x -= 2
	fmt.Println("x -= 2 ->", x) // ผลลัพธ์: 6

	// คูณค่าแล้วกำหนดค่า
	x *= 4
	fmt.Println("x *= 4 ->", x) // ผลลัพธ์: 24

	// หารค่าแล้วกำหนดค่า
	x /= 2
	fmt.Println("x /= 2 ->", x) // ผลลัพธ์: 12

	// หารเอาเศษแล้วกำหนดค่า
	x %= 5
	fmt.Println("x %= 5 ->", x) // ผลลัพธ์: 2

	// AND บิตแล้วกำหนดค่า
	x &= 3
	fmt.Println("x &= 3 ->", x) // ผลลัพธ์: 2

	// OR บิตแล้วกำหนดค่า
	x |= 1
	fmt.Println("x |= 1 ->", x) // ผลลัพธ์: 3

	// XOR บิตแล้วกำหนดค่า
	x ^= 1
	fmt.Println("x ^= 1 ->", x) // ผลลัพธ์: 2

	// เลื่อนบิตซ้ายแล้วกำหนดค่า
	x <<= 1
	fmt.Println("x <<= 1 ->", x) // ผลลัพธ์: 4

	// เลื่อนบิตขวาแล้วกำหนดค่า
	x >>= 1
	fmt.Println("x >>= 1 ->", x) // ผลลัพธ์: 2
}
