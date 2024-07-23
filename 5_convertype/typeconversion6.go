// แปลงจากบูลีนเป็นจำนวนเต็ม (แสดงในรูปแบบ integer เพราะ Go ไม่รองรับการแปลงจาก bool เป็น int โดยตรง)

package main

import "fmt"

func TypeConversion6() {
	var b bool = true
	var i int
	if b {
		i = 1
	} else {
		i = 0
	}
	fmt.Printf("b: %t, i: %d\n", b, i)
	// ผลลัพธ์: b: true, i: 1
}
