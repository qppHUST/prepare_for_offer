package main

import (
	"fmt"
	"time"
)

type Task struct {
	A int
	B int
}

var bucket = make(chan Task, 100)

func producer() {
	for i := 0; i < 10; i++ {
		bucket <- Task{A: i, B: i + 1}
	}
}

func consumer() {
	for i := 0; i < 10; i++ {
		task := <-bucket
		fmt.Printf("%d %d\n", task.A, task.B)
	}
}

func main() {
	go producer()
	go consumer()
	time.Sleep(time.Second * 10)
}
