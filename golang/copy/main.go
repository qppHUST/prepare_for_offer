package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := []int{}
	copy(b, a)
	fmt.Printf("%p  %p", a, b)
}
