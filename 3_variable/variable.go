package main

import "fmt"

var myworldGlob string = "Awesome Global Variable"

func main() {

	// Variable in Go (ตัวแปรใน Go)
	// กฏการตั้งชื่อตัวแปร
	// 1. ชื่อตัวแปรต้องเริ่มต้นด้วยตัวอักษรหรือ _ (underscore)
	// 2. ชื่อตัวแปรสามารถประกอบไปด้วยตัวอักษร a-z, A-Z, ตัวเลข 0-9 และ _ (underscore)
	// 3. ชื่อตัวแปรเป็น case-sensitive
	// 4. ชื่อตัวแปรไม่สามารถเป็น keyword หรือ reserved word ของภาษา Go

	// ตัวอย่างคำสงวนของภาษา Go
	// break        default      func         interface    select
	// case         defer        go           map          struct
	// chan         else         goto         package      switch
	// const        fallthrough  if           range        type
	// continue     for          import       return       var

	// 1. Variable declaration
	// var <variable_name> <data_type>
	var str string
	var (
		num   float32
		count int
	) // ประกาศตัวแปรหลายตัวในบรรทัดเดียวกัน
	var Pi float32

	// 2. Variable assignment
	str = "Go Programming Language"
	num = 100.003
	count = 10
	Pi = 3.14

	fmt.Println(str)
	fmt.Println(num)
	fmt.Println(count)
	fmt.Println(Pi)

	fmt.Println("--------------------")

	// 3. Variable declaration and assignment
	// var <variable_name> <data_type> = <value>
	var name string = "John Doe"
	var age int = 30
	var weight float32 = 70.5

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(weight)

	fmt.Println("--------------------")

	// 4. Variable declaration and assignment with type inference
	// <variable_name> := <value>
	// ใช้เมื่อต้องการประกาศตัวแปรและกำหนดค่าในครั้งเดียวกัน
	// ไม่จำเป็นต้องระบุ data type
	// ใช้เมื่อต้องการประกาศตัวแปรและกำหนดค่าในครั้งเดียวกัน
	// ไม่จำเป็นต้องระบุ data type

	// ตัวอย่างการใช้ type inference
	var a = 10
	var b = 10.5
	var c = "Hello, World!"
	var d = true
	var num1, num2, str1, str2 = 100, 200, "Hello", "World"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(num1)
	fmt.Println(num2)
	fmt.Println(str1)
	fmt.Println(str2)

	fmt.Println("--------------------")

	// 5. short variable declaration
	// <variable_name> := <value>
	product := "Computer"
	price := 10000
	quantity := 10

	fmt.Println(product)
	fmt.Println(price)
	fmt.Println(quantity)

	fmt.Println("--------------------")

	// 6. Multiple short variable declaration
	// <variable_name1>, <variable_name2> := <value1>, <value2>
	var x, y, z = 10, 20, 30
	var i, j, k = 10.5, 20.5, 30.5
	var p, q, r = "Hello", "World", "Go"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(p)
	fmt.Println(q)
	fmt.Println(r)

	fmt.Println("--------------------")

	// 7. Constant
	// ประกาศค่าคงที่
	// const <constant_name> <data_type> = <value>
	const pi float32 = 3.14
	const vat = 7.0
	const name1, name2 = "John", "Doe"
	// vat = 10.0 // error: cannot assign to vat

	fmt.Println(pi)
	fmt.Println(vat)
	fmt.Println(name1)
	fmt.Println(name2)

	fmt.Println("--------------------")

	// 8. Variable scope
	// ตัวแปรที่ประกาศใน function จะมีขอบเขตใน function นั้นๆ
	// ตัวแปรที่ประกาศใน block จะมีขอบเขตใน block นั้นๆ
	// ตัวแปรที่ประกาศด้วย var จะมีขอบเขตใน function นั้นๆ
	// ตัวแปรที่ประกาศด้วย := จะมีขอบเขตใน block นั้นๆ

	// ตัวอย่างขอบเขตของตัวแปร
	var mywordLocal string = "Awesome Local Variable"
	fmt.Println(myworldGlob) // แสดงค่าตัวแปร myword ที่ประกาศใน package main
	fmt.Println(mywordLocal) // แสดงค่าตัวแปร myword ที่ประกาศใน function main

	// เรียกใช้ฟังก์ชัน test
	test()
}

func test() {
	fmt.Println(myworldGlob) // แสดงค่าตัวแปร myword ที่ประกาศใน package main
	// fmt.Println(mywordLocal) // error: undefined: mywordLocal
}
