package main

import "fmt"

func main() {
	list := make([]int, 0, 5)
	changeSlice(list)
	fmt.Println(list)
}

func changeSlice(list []int) {
	// list = append(list, 1, 2, 3, 4, 5, 6, 7, 8, 9, 7, 8, 5, 4, 5, 5, 5, 5, 6, 3, 6)
	list = append(list, 2)
}
