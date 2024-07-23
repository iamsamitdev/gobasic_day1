// แปลงจากจำนวนทศนิยมเป็นจำนวนเต็ม

package main

import (
	"fmt"
)

func TypeConversion2() {
	var f float64 = 42.56
	var i int = int(f)
	fmt.Printf("f: %f, i: %d\n", f, i)
	// ผลลัพธ์: f: 42.560000, i: 42
}
