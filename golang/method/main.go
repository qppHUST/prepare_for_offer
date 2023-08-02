package main

import (
	"fmt"
)

type animal struct {
	name string
}

func change(animal animal) {
	animal.name = "lala"
}

func main() {
	animal1 := animal{name: "heihei"}
	change(animal1)
	fmt.Println(animal1)
}
