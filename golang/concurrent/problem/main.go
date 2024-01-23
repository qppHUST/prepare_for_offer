package main

import (
	"fmt"
)

// 三个线程交替打印ABC，非循环
var str = "ABC"

func print(tunnel chan int, done chan struct{}) {
	for {
		index := <-tunnel
		fmt.Println(string(str[index]))
		done <- struct{}{}
	}
}

func main() {
	tuneel := make(chan int, 1)
	done := make(chan struct{}, 1)
	go print(tuneel, done)
	go print(tuneel, done)
	go print(tuneel, done)
	done <- struct{}{}
	for i := 0; i < 10; i++ {
		<-done
		tuneel <- i % 3
	}
}
