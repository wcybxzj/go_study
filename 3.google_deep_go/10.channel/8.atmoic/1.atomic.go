package main

import (
	"fmt"
	"sync"
	"time"
)

//atomicInt
type atomicInt int

func (a *atomicInt) increment() {
	*a++
}

func (a *atomicInt) get() int {
	return int(*a)
}

//atomicInt2
type atomicInt2 struct {
	value int
	lock  sync.Mutex
}

func (a *atomicInt2) increment() {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

//第一次感受闭包的巧妙:
//通过匿名函数+defer来实现一块代码的原子性
//原理:让defer在一个匿名函数中,匿名函数结束defer执行
func (a *atomicInt2) increment3() {
	fmt.Println("safe increment")
	//其他代码
	//其他代码
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
	//其他代码
	//其他代码
}

func (a *atomicInt2) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return a.value
}

//测试1:
//go run -race atomic.go 说有冲突
func test1() {
	var a atomicInt
	a.increment()

	go func() {
		a.increment() //并发冲突点
	}()

	time.Sleep(time.Millisecond) //父协程灯1ms让子协程有机会执行
	fmt.Println(a)               //并发冲突点
}

//测试2:
//用传统并发实现一个协程安全的原子化的int
//开发中尽量使用:go本身就有 atomic.AddInt32()来实现协程原子化的Add
//go run -race atomic.go 无冲突
func test2() {
	var a atomicInt2
	a.increment()

	go func() {
		a.increment() //并发冲突点
	}()

	time.Sleep(time.Millisecond) //父协程灯1ms让子协程有机会执行
	fmt.Println(a.get())         //并发冲突点
}

//测试3:
//通过匿名函数+defer来实现一块代码的原子性
//go run -race atomic.go 无冲突
func test3() {
	var a atomicInt2
	a.increment3()

	go func() {
		a.increment3() //并发冲突点
	}()

	time.Sleep(time.Millisecond) //父协程灯1ms让子协程有机会执行
	fmt.Println(a.get())         //并发冲突点
}

func main() {
	//test1()
	//test2()
	test3()
}
