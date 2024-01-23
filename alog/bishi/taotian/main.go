package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	var parseList []string
	str = str[:len(str)-2]
	for i := 0; i < len(str); i++ {
		if str[i] == '(' || str[i] == ')' {
			parseList = append(parseList, string(str[i]))
		} else if str[i] >= 65 && str[i] <= 90 {
			//是大写字母，开始找单词
			var builder strings.Builder
			builder.WriteString(string(str[i]))
			index := i
			for {
				index++
				if str[index] >= 97 && str[index] <= 122 || str[index] >= 48 && str[index] <= 57 {
					//小写字母或者数字
					builder.WriteString(string(str[index]))
				} else {
					//大写字母，出去
					break
				}
			}
			parseList = append(parseList, builder.String())
			i = index - 1
		} else if str[i] >= 48 && str[i] <= 57 {
			var builder strings.Builder
			builder.WriteString(string(str[i]))
			index := i
			for {
				index++
				if index < len(str) && str[index] >= 48 && str[index] <= 57 {
					builder.WriteString(string(str[index]))
				} else {
					break
				}
			}
			i = index - 1
			parseList = append(parseList, builder.String())
		}
	}
	var stack []string
	for _, item := range parseList {
		if item == "(" {
			stack = append(stack, item)
		} else if item == ")" {

		} else {
			stack = append(stack, item)
		}
	}
}
