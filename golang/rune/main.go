package main

import (
	"fmt"
)

func main() {
	str := "qpphust邱攀攀"
	str = string([]rune(str)[:8])
	fmt.Println(str)
}
