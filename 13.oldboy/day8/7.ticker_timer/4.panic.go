package main

import (
	"fmt"
	"runtime"
	"time"
)

//运行时panic:assignment to entry in nil map
//如果不进行recover整个进程就挂了
//加了recover,其他的协程不受影响
func test() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic:", err)
		}
	}()

	var m map[string]int
	m["stu"] = 100
}

func work() {
	for {
		fmt.Println("working")
		time.Sleep(time.Second)
	}
}

func main() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)

	go test()
	for i := 0; i < 100; i++ {
		go work()
	}

	for {
		fmt.Println("main is alive")
		time.Sleep(time.Second * 1)
	}
}
