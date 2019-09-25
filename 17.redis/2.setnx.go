package main

import (
	"fmt"
	"go_study/17.redis/model"
	"time"
)

func test() {
	set, err := model.Client.SetNX("key", "value", 10*time.Second).Result()

	if err != nil {
		panic("err:%v"+ err.Error())
	}

	fmt.Println(set)

	if set == true {
		fmt.Println(" fail", "key" )
	} else {
		fmt.Println("set success, key:%s", "key" )
	}

}

func main() {
	model.ExampleNewClient()
	test()
}