package main

import "fmt"

//channel应该被用来通知
// buffered channel 会产生两次复制
// unbuffered channel 只会产生一次复制
// unbuffered channel 可以保证receive之后 send才返回
func main() {
	chan1 := make(chan int)
	send(chan1)
	receive(chan1)
	close(chan1)
	i := <-chan1
	chan1 <- 2
	fmt.Println(i)
}

func send(chan1 chan<- int) {

}

func receive(chan2 <-chan int) {

}
