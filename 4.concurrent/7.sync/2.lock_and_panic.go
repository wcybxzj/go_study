package main

import (
	"fmt"
	"sync"
	"time"
)

type rpcSt2 struct {
	Host string
	rwLock sync.RWMutex
}

var rpc2 rpcSt2

func func1() {
	rpc2.rwLock.Lock()
	fmt.Println("func1 lock")
	panic("111")
	rpc2.rwLock.Unlock()
}

func func2() {
	rpc2.rwLock.Lock()
	fmt.Println("func2 lock")
	rpc2.rwLock.Unlock()
}

func test1() {
	go func1()
	time.Sleep(time.Second*2)
	go func2()
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		for  {
			time.Sleep(time.Second)
			fmt.Println("222")
		}
	}()

	panic("1111")

	//test1()

	time.Sleep(time.Second*10)
}
