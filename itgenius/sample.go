package sample

import (
	"fmt"
	personal "github.com/iamsamitdev/gobasic/itgenius/internal"
)

func HelloITGenius() {
	fmt.Println("Hello, from ITGenius!")
	personal.HelloPersonal()
}

func Add(a, b int) int {
	return a + b
}
