package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	str = str[:len(str)-2]
	stack := []rune{}
	ans := 0
	for _, item := range str {
		if len(stack) == 0 {
			stack = append(stack, item)
			continue
		}
		stackLastRune := stack[len(stack)-1]
		if stackLastRune == '(' && item == ')' ||
			stackLastRune == '[' && item == ']' {
			stack = stack[:len(stack)-1]
		} else if stackLastRune == ')' || stackLastRune == ']' {
			stack = stack[:len(stack)-1]
			ans++
		} else {
			stack = append(stack, item)
		}
	}
	ans = ans + len(stack)
	fmt.Println(ans)
}
