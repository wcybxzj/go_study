package main

import (
	"fmt"
	"go_study/17.redis/model"
)

import "github.com/go-redis/redis"



func ExampleClient() {
	err := model.Client.Set("key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := model.Client.Get("key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := model.Client.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func main() {
	model.ExampleNewClient()
	ExampleClient()
}