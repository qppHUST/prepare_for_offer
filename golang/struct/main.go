package main

import (
	"fmt"
	"unsafe"
)

type animal struct{}

type a struct {
	name string
}

type b struct {
	a
}

func (animal *animal) move() {
	fmt.Println("11")
}

func main() {
	var animal animal
	fmt.Println(unsafe.Sizeof(animal))

	var b b
	fmt.Println(b.name)
}
