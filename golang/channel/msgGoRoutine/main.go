package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string    // 真正要传输的数据
	wait chan bool //
}

func fanIn(inputs ...<-chan Message) <-chan Message {
	c := make(chan Message)
	for i := range inputs {
		input := inputs[i]
		go func() {
			for {
				c <- <-input
			}
		}()
	}
	return c
}

// `boring` 是一个返回管道的方法，该管道用于和 `boring` 方法通信
func boring(msg string) <-chan Message {
	c := make(chan Message)
	waitForIt := make(chan bool)
	go func() {
		for i := 0; ; i++ {
			c <- Message{
				str:  fmt.Sprintf("%s %d", msg, i),
				wait: waitForIt, // 将管道注入到返回值中，用于协程之间的通信
			}
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)

			// 协程需要在这里等到接收到信息才能够继续执行后面的逻辑
			<-waitForIt
		}

	}()
	return c
}

func main() {
	// merge 2 channels into 1 channel
	c := fanIn(boring("Joe"), boring("Ahn"))

	for i := 0; i < 5; i++ {
		msg1 := <-c // 等到从管道中接受数据
		fmt.Println(msg1.str)
		msg2 := <-c
		fmt.Println(msg2.str)

		// 由于 boring 协程中需要等待 wait 信号才能继续执行，所以这一步能够保证两个协程都能够输出一次数据
		msg1.wait <- true // main 协程允许 boring 协程继续执行任务
		msg2.wait <- true // main 协程允许 boring 协程继续执行任务
	}
	fmt.Println("You're both boring. I'm leaving")
}
