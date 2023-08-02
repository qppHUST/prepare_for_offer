package main

import (
	"fmt"
	"math/rand"
	"time"
)

// `boring` 是一个返回管道的方法，该管道用于和 `boring` 方法通信
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
		}

	}()
	return c
}

func main() {
	c := boring("Joe")

	// timeout 的类型为：`<-chan Time`
	timeout := time.After(5 * time.Second) // 这个方法会在 5 秒钟向 timeout 管道写入数据
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeout:
			fmt.Println("You talk too much.")
			return
		}
	}
}
