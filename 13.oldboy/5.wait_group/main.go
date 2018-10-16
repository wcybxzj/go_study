package main

import (
	"fmt"
	"sync"
	"time"
)

func calc1(w *sync.WaitGroup, i int) {
	w.Add(1)
	fmt.Println("calc:", i)
	time.Sleep(time.Second)
	w.Done()
}

//测试1:失败,一个任务也不能执行
//
//输出:
//all goroutine finish
//
//分析:
//因为可能一个worker还没执行,所以WaitGroup为0,主协程直接退出
func test1() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go calc1(&wg, i)
	}

	wg.Wait()
	fmt.Println("all goroutine finish")
}

func calc2(w *sync.WaitGroup, i int) {
	fmt.Println("calc:", i)
	time.Sleep(time.Second)
	w.Done()
}

//测试2: 正确
func test2() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go calc2(&wg, i)
	}

	wg.Wait()
	fmt.Println("all goroutine finish")
}

func main() {
	test1()
	//test2()
}
