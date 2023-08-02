package main

import "fmt"

type animal struct {
	name string
}

func main() {
	a1 := &animal{"puu"}
	a2 := *a1
	a2.name = "biu"
	fmt.Println(*a1)
	fmt.Println(a2)

}
