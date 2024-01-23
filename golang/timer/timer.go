package main

import (
	"fmt"
	"time"
)

func main() {
	item := new(task)
	fmt.Printf(item.name)
}

type task struct {
	name        string
	f           func()
	executeTime time.Time
}

type delayQueue struct {
	taskHeap []*task //min heap
}

func (obj *delayQueue) format() {
	cursor := len(obj.taskHeap) - 1
	for {
		father := -1
		if cursor%2 == 0 {
			father = cursor / 2
		} else {
			father = (cursor - 1) / 2
		}
		if obj.taskHeap[cursor].executeTime.Before(obj.taskHeap[father].executeTime) {
			obj.taskHeap[cursor], obj.taskHeap[father] = obj.taskHeap[father], obj.taskHeap[cursor]
		} else {
			break
		}
	}
}

func (obj *delayQueue) put(item *task) {
	obj.taskHeap = append(obj.taskHeap, item)
	obj.format()
}

func start() {

}
