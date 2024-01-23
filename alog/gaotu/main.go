package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	readString = readString[:len(readString)-2]
	dp := make([]int, len(readString))
	dp[0] = 1
	max := math.MinInt64
	for i := 1; i < len(readString); i++ {
		ok := false
		for j := i - 1; j >= 0; j-- {
			if readString[j] == readString[i] {
				dp[i] = i - j
				ok = true
				break
			}
		}
		if !ok {
			dp[i] = dp[i-1] + 1
		}
		if dp[i] > max {
			max = dp[i]
		}
	}
	fmt.Println(max)
}
