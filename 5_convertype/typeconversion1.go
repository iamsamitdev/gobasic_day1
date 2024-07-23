// แปลงจากจำนวนเต็มเป็นจำนวนทศนิยม
package main

import (
	"fmt"
)

func TypeConversion1() {
	var i int = 42
	var f float64 = float64(i)
	fmt.Printf("i: %d, f: %f\n", i, f)
	// ผลลัพธ์: i: 42, f: 42.000000
}
