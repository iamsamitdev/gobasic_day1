// แปลงจากสตริงเป็นจำนวนเต็ม

package main

import (
	"fmt"
	"strconv"
)

func TypeConversion4() {
	var s string = "42"
	var i int
	var err error
	i, err = strconv.Atoi(s)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("s: %s, i: %d\n", s, i)
	// ผลลัพธ์: s: 42, i: 42
}
