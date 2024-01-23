package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	var list []int
	fmt.Println(list)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	select {
	case <-ctx.Done():
		break

	}
}
