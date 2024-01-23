package main

import (
	"sync"
	"time"
)

type TokenBucketlimiter struct{
	token int64
	standard int64
	lock sync.Mutex
	putRate int64
	lastTime int64
}

func (obj *TokenBucketlimiter)request()bool{
	obj.lock.Lock()
	defer obj.lock.Unlock()
	if obj.token == 0{
		return false
	}
	obj.token--
	nowTime := time.Now().Unix()
	toPut := obj.putRate*(nowTime-obj.lastTime)
	obj.token = (obj.token+toPut)%obj.standard
	return true
}

func (obj *TokenBucketlimiter)New(){
	go func() {

	}()
}

func main() {

}
