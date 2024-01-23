package main

func main() {
	q := make(chan int)
	<-q
}
