package main

import "fmt"

func exampleCompareOperator() {
	var x bool

	// เท่ากับ
	x = (5 == 3)
	fmt.Println("5 == 3 =", x) // ผลลัพธ์: false

	// ไม่เท่ากับ
	x = (5 != 3)
	fmt.Println("5 != 3 =", x) // ผลลัพธ์: true

	// มากกว่า
	x = (5 > 3)
	fmt.Println("5 > 3 =", x) // ผลลัพธ์: true

	// น้อยกว่า
	x = (5 < 3)
	fmt.Println("5 < 3 =", x) // ผลลัพธ์: false

	// มากกว่าหรือเท่ากับ
	x = (5 >= 3)
	fmt.Println("5 >= 3 =", x) // ผลลัพธ์: true

	// น้อยกว่าหรือเท่ากับ
	x = (5 <= 3)
	fmt.Println("5 <= 3 =", x) // ผลลัพธ์: false
}
