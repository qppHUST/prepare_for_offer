package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewReader(os.Stdin)
	var list []int
	var builder strings.Builder
	numStr, err := scanner.ReadString('\n')
	if err != nil {
		return
	}
	numStr = numStr[:len(numStr)-2]
	num, _ := strconv.Atoi(numStr)
	for num > 0 {
		num--
		str, err := scanner.ReadString('\n')
		str = str[:len(str)-2]
		if err != nil {
			return
		}
		for _, item := range strings.Split(str, ",") {
			numToStore, _ := strconv.Atoi(item)
			list = append(list, numToStore)
		}
	}
	sort.Ints(list)
	for _, item := range list {
		builder.WriteString(strconv.Itoa(item) + ",")
	}
	toPrint := builder.String()
	if len(toPrint) == 0 {
		fmt.Println("")
	} else {
		fmt.Println(toPrint[:len(toPrint)-1])
	}
}
