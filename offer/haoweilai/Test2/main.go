package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	var builder strings.Builder
	str, _ := scanner.ReadString('\n')
	str = str[:len(str)-2]
	flag := true
	if str[0] == '-' {
		str = str[1:]
		flag = false
	}
	list := strings.Split(str, "")
	for i := len(list) - 1; i >= 0; i-- {
		builder.WriteString(list[i])
	}
	atoi, err := strconv.Atoi(builder.String())
	if err != nil {
		return
	}
	if !flag {
		atoi = -atoi
	}
	if atoi > math.MaxInt32 || atoi < math.MinInt32 {
		fmt.Println(0)
	} else {
		fmt.Println(atoi)
	}
}
