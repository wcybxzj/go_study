package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func doWork(id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		wg.Done() //删除一次任务
	}
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		wg: wg,
	}
	go doWork(id, w.in, w.wg)
	return w
}

func test_1_1() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	//写法1:(正确)
	wg.Add(20) //加20个任务(必须怎么写)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}
	//等待所有任务结束,只有任务为0才解除阻塞
	wg.Wait()
}

func test_1_2() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		wg.Add(1) //写法2:正常
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		wg.Add(1) //写法2:正常
		worker.in <- 'A' + i
	}
	//等待所有任务结束,只有任务为0才解除阻塞
	wg.Wait()
}

//报错:
//panic: sync: negative WaitGroup counter
//意思:WaitGroup counter 出错为负数

//分析:
//main go add前,worker go 已经消费完并且Done,造成WaitGroup counter为负
func test_1_3() {
	var wg sync.WaitGroup
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		wg.Add(1) //每次加一个任务(Bug)
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		wg.Add(1) //每次加一个任务(Bug)
	}
	//等待所有任务结束,只有任务为0才解除阻塞
	wg.Wait()
}

//测试1:(成功)
//解决2.done.go中test2()死锁的第三种方法:
//用Golang版本的WaitGroup, 让主协程来等待任务结束
//Add是添加任务, Done是完成了任务, Wait是等待任务
func test1() {
	test_1_1()//成功
	//test_1_2()//成功
	//test_1_3() //失败
}

type worker2 struct {
	in   chan int
	done func()
}

func doWork2(id int, w worker2) {
	for n := range w.in {
		fmt.Printf("worker %d received %c\n", id, n)
		w.done()
	}
}

func createWorker2(id int, wg *sync.WaitGroup) worker2 {
	w := worker2{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork2(id, w)
	return w
}

//测试2:(成功)
//对test1()进行包装
func test2() {
	var wg sync.WaitGroup
	var workers [10]worker2
	for i := 0; i < 10; i++ {
		workers[i] = createWorker2(i, &wg)
	}

	wg.Add(20) //加20个任务(必须怎么写)
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	//等待所有任务结束,只有任务为0才解除阻塞
	wg.Wait()
}

func main() {
	test1()
	//test2()
}
