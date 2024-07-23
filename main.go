package main

import (
	"fmt"
	"github.com/google/uuid"
	sample "github.com/iamsamitdev/gobasic/itgenius"
)

func main() {
	fmt.Println("Hello, World!")
	sample.HelloITGenius()
	fmt.Println(sample.Add(10, 20))

	// Generate a new UUID
	id := uuid.New()
	fmt.Println("UUID: ", id)
}
