package main

import (
	"fmt"
	"sync"
	"time"
)

//test1:不使用锁并发下map直接报错
//报错：
//fatal error: concurrent map read and map write
func test1() {
	m := make(map[int]int)
	go func() {
		for {
			_ = m[1]
		}
	}()
	go func() {
		for {
			m[2] = 2
		}
	}()
	select {}
}

/*
test2:使用手动给map加锁
成功
go run 2.map.go --race
some_key: 0
some_key: 123
*/
func test2() {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	go func() {
		counter.RLock()
		n := counter.m["some_key"]
		counter.RUnlock()
		fmt.Println("some_key:", n)
	}()

	go func() {
		counter.Lock()
		counter.m["some_key"] = 123
		n := counter.m["some_key"]
		counter.Unlock()
		fmt.Println("some_key:", n)
	}()

	time.Sleep(time.Second * 1)
}

func main() {
	test1()
	//test2()
}
