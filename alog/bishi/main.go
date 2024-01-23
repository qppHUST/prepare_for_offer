package main

import (
	"bufio"
	"fmt"
	"os"
)

//给你一个字符串 s ，请你反转字符串中 单词 的顺序。
//单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
//返回 单词 顺序颠倒且 单词 之间用单个空格连接的结果字符串。
//注意：输入字符串 s中可能会存在前导空格、尾随空格或者单词间的多个空格。
//返回的结果字符串中，单词间应当仅用单个空格分隔，且不包含任何额外的空格。

//"a good   example     " => "example good a"

func main() {
	reader := bufio.NewReader(os.Stdin)
	readString, _ := reader.ReadString('\n')
	readString = readString[:len(readString)-2]
	fmt.Println(readString)
	ans := method(readString)
	fmt.Println("----------------------------------")
	fmt.Println(ans)
}

func method(str string) string {
	cursor := len(str) - 1
	builder := ""
	ans := ""
	begin := false
	for cursor > -1 {
		begin = true
		item := str[cursor]
		if item != ' ' {
			itoa := string(item)
			builder = itoa + builder
			if cursor == 0 {
				ans = ans + builder
			}
		} else {
			if begin {
				ans = ans + builder + " "
				builder = ""
			}
			begin = false
		}
		cursor--
	}
	return ans
}
