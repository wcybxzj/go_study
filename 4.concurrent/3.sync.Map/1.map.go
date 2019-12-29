package main

import (
	"fmt"
	"os"
	"strconv"
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

/*
说明当rlock上锁的时候,独占的lock就只能阻塞
输出:
rlocking i: 0
rlocking i: 1
rlocking i: 2
some_key: 0
some_key: 123
*/
func test3() {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	go func() {
		counter.RLock()
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			fmt.Println("rlocking i:", i)
		}
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

	time.Sleep(time.Second * 100)
}

//更真实的读写情况
func test4() {
	var counter = struct {
		sync.RWMutex
		m map[string]int
	}{m: make(map[string]int)}

	for i:=0;i<1000 ;i++  {
		//reader 1000 goroutine
		go func() {
			for {
				counter.RLock()
				fmt.Println("read-read-read-read i:"+strconv.Itoa(i))
				for i := 0; i < 3; i++ {
					time.Sleep(time.Second)
					fmt.Println("rlocking i:", i)
				}
				n := counter.m["some_key"]
				counter.RUnlock()
				time.Sleep(time.Second)
				fmt.Println("some_key:", n)
			}
		}()
	}


	//write 1 goroutine
 	go func() {
		counter.Lock()
		fmt.Println("write-write-write-write")
		counter.m["some_key"] = 123
		n := counter.m["some_key"]
		counter.Unlock()
		os.Exit(0)
		fmt.Println("some_key:", n)
	}()

	time.Sleep(time.Second * 100)
}

func main() {
	//test1()
	//test2()
	//test3()
	test4()
}
