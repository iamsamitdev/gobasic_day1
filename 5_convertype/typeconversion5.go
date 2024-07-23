// แปลงจากจำนวนเต็มเป็นบูลีน (แสดงในรูปแบบ integer เพราะ Go ไม่รองรับการแปลงจาก int เป็น bool โดยตรง)

package main

import "fmt"

func TypeConversion5() {
	var i int = 1
	var b bool = (i != 0)
	fmt.Printf("i: %d, b: %t\n", i, b)
	// ผลลัพธ์: i: 1, b: true
}
