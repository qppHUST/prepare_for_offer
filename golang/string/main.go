package main

import "fmt"

func main() {
	str := "邱攀攀qpphust"
	byteData := []byte(str)
	// byteData[0] = 'a'
	fmt.Println(string(byteData))

}
