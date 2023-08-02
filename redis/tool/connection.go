package tool

import (
	"log"

	"github.com/go-redis/redis"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     "localhost:8999",
		Password: "qpphust",
		DB:       0,
	})
	_, error := client.Ping().Result()
	if error != nil {
		log.Fatal(error)
	}
}

func GetConnection() *redis.Client {
	return client
}

func RunCmd(args ...interface{}) *redis.Cmd {
	return client.Do(args...)
}
