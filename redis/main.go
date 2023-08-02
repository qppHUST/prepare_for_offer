package main

import (
	"fmt"
	"log"

	"github.com/qppHUST/prepareForOffer/redis/tool"
)

func main() {
	err := tool.RunCmd("set", "a", "1000").Err()
	if err != nil {
		log.Fatal(err)
	}
	ans, error := tool.RunCmd("get", "a").Result()
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println(ans)
}

// func a() <-chan struct{} {
// 	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
// 	if someError {
// 		cancel()
// 	}
// 	return ctx.Done()
// }
