package main

import "fmt"

func main() {
	num := 0
	ans := 0
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}
	notZhiShu := func(num int) bool {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				return true
			}
		}
		return false
	}
	for i := 2; i < num; i++ {
		if !notZhiShu(i) {
			ans++
		}
	}
	fmt.Println(ans)
}
