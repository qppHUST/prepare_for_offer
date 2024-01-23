package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var list []int32
	var lock sync.Mutex
	chanelToAdvance := make(chan struct{})
	chanelToFinish := make(chan struct{})
	index := 0
	str1 := "sasasa"
	str2 := "lslskncj"
	go func(str string) {
		ticker := time.NewTicker(time.Second)
		for _, c := range str {
			<-ticker.C
			lock.Lock()
			list = append(list, c)
			lock.Unlock()
			chanelToAdvance <- struct{}{}
		}
	}(str1)
	go func(str string) {
		ticker := time.NewTicker(time.Second)
		for _, c := range str {
			<-ticker.C
			lock.Lock()
			list = append(list, c)
			lock.Unlock()
			chanelToAdvance <- struct{}{}
		}
	}(str2)
	go func() {
		for {
			<-chanelToAdvance
			fmt.Print(string(list[index]))
			index++
			if index == len(str1)+len(str2) {
				chanelToFinish <- struct{}{}
			}
		}
	}()
	<-chanelToFinish
}
