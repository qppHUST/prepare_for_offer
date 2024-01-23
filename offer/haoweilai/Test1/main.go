package main

import (
	"bufio"
	"os"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	str1, _ := scanner.ReadString('\n')
	str1 = str1[:len(str1)-2]
	str2, _ := scanner.ReadString('\n')
	str2 = str2[:len(str2)-2]
}
