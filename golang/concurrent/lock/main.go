package main

import (
	"runtime"
	"sync"
)

var lock sync.RWMutex

func main() {
	lock.Lock()
	runtime.Gosched()
}
