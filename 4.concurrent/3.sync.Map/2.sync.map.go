package main

import (
	"fmt"
	"sync"
	"time"
)

func test3() {
	var counters sync.Map

	//read
	go func() {
		for {
			v, ok := counters.Load("habr")
			if ok {
				fmt.Println(v)
			}

		}
	}()

	//write
	go func() {
		for {
			counters.Store("habr", 42)
		}
	}()

	//delete
	go func() {
		for {
			counters.Delete("habr")
		}
	}()

	select {}
}

/*
LoadOrStore():的意思是key存在就读取，不存在就保存新值
输出:
42
42
*/
func LoadOrStore4() {
	var counters sync.Map

	counters.Store("habr", 42)

	go func() {
		for {
			time.Sleep(time.Second)
			v2, ok := counters.LoadOrStore("habr", 13)
			if ok {
				fmt.Println(v2)
			} else {
				fmt.Println("LoadOrStore() fail!")
			}

		}
	}()
	select {}
}

/*
LoadOrStore():的意思是key存在就读取，不存在就保存新值
输出:
LoadOrStore() fail!
13
13
13
*/
func LoadOrStore5() {
	var counters sync.Map

	counters.Store("habr", 42)

	go func() {
		for {
			time.Sleep(time.Second)
			v2, ok := counters.LoadOrStore("habr2", 13)
			if ok {
				fmt.Println(v2)
			} else {
				fmt.Println("LoadOrStore() fail!")
			}
		}
	}()
	select {}
}

func test6() {

	var counters sync.Map

	go func() {
		counters.Store("habr", 42)
	}()

	go func() {
		for {
			counters.Range(func(k, v interface{}) bool {
				fmt.Println("key:", k, ", val:", v)
				return true // if false, Range stops
			})
			time.Sleep(time.Second)
		}
	}()

	select {}
}

//解决1.map.go中test4()的问题
func test7() {
	var m sync.Map


	m.Store("")

	go func() {
		for {
			for i := 0; i < 3; i++ {
				time.Sleep(time.Second)
				fmt.Println("rlocking i:", i)
			}
			fmt.Println("some_key:", m[""])
		}
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

func main() {
	//test3()
	//LoadOrStore4()
	//LoadOrStore5()
	test6()
}
