package main

import (
	"fmt"
)

type dog interface {
	Move()
}

func main() {
	var kk dog
	kk.Move()
	switch kk.(type) {
	case dog:
		fmt.Println("dog")
	}
}
