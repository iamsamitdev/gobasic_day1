package main

import (
	"fmt"
	"html"
)

// Glabal variable
var xml = "outside-myxml"

func main() {

	// Local variable
	xml := "inside-myxml"
	fmt.Println(xml) // EscapeString:  myxml

	// Data type in Go
	// 2 types of data type in Go
	// 1. Primitive data type (basic data type)
	// 2. Derived data type (composite data type)

	// Numeric data type
	// - Integer data type (จำนวนเต็ม)
	var a int8 = 10     // range: -128 to 127
	var b int16 = 20    // range: -32768 to 32767
	var d int64 = 40    // range: -9223372036854775808
	var e uint8 = 50    // range: 0 to 255
	var f uint16 = 60   // range: 0 to 65535
	var g uint32 = 70   // range: 0 to 4294967295
	var h uint64 = 80   // range: 0 to 18446744073709551615
	var i int = 90      // int is alias of int32 or int64
	var j uint = 100    // uint is alias of uint32 or uint64
	var k uintptr = 110 // uintptr is an unsigned integer type that is large enough to store the uninterpreted bits of a pointer value
	var l byte = 120    // byte is alias of uint8
	var m rune = 130    // rune is alias of int32

	// - Floating-point data type (จำนวนทศนิยม)
	var n float32 = 3.14
	var o float64 = 3.141592653589793

	// - Complex data type (จำนวนเชิงซ้อน)
	var p complex64 = 3 + 4i
	var q complex128 = 5 + 6i

	// Boolean data type
	var r bool = true
	var s bool = false

	// String data type (ข้อความ)
	var t string = "Hello, World!" // string literal
	var u string = `Hello,
	World!` // raw string literal

	// Escape sequence (อักขระพิเศษ)
	// \a  Bell (alert) (เสียงเตือน)
	// \b  Backspace (ลบเครื่องหมาย)
	// \f  Form feed (ขึ้นบรรทัดใหม่)
	// \n  New line (ขึ้นบรรทัดใหม่)
	// \r  Carriage return (เลื่อนเคอร์เซอร์ไปตำแหน่งแรกของบรรทัด)
	// \t  Horizontal tab (เว้นวรรคแนวนอน)
	// \v  Vertical tab (เว้นวรรคแนวตั้ง)
	// \'  Single quote (เครื่องหมายอัญประสงค์เดี่ยว)
	// \"  Double quote (เครื่องหมายอัญประสงค์คู่)
	// \\  Backslash (เครื่องหมายเบ็ดเสร็จ)

	println(a, b, d, e, f, g, h, i, j, k, l, m, n, o, p, q, r, s, t, u)

	const str = `<a href="mailto:email@example.com">
				Send Email</a>` // raw string literal
	en := html.EscapeString(str)  // escape string
	de := html.UnescapeString(en) // unescape string

	fmt.Println("Encoded: ", en)
	fmt.Println("Decoded: ", de)

}
