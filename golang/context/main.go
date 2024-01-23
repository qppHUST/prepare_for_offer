package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	select {
	case <-ctx.Done():
		fmt.Println("done")
	default:
	}
}
