package model

import (
	"fmt"
	"github.com/go-redis/redis"
)

var Client *redis.Client

func ExampleNewClient() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "nqyong_redis_123", // no password set
		DB:       0,  // use default DB
	})

	pong, err := Client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}