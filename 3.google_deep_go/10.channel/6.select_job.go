package main

import (
	"fmt"
	"math/rand"
	"time"
)

//在1ms->1.5秒内产生1个任务
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) *
				time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

//每1秒消耗一个任务
func worker2(id int, c chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker2(id int) chan<- int {
	c := make(chan int)
	go worker2(id, c)
	return c
}

//测试6:(失败)
//残留问题:
//假如generator()每1ms->1.5秒产生一个数据，而worker每1秒才消耗一次数据
//generator()生成的任务会有大量丢失
/*
输出:
worker 0 received 0
worker 0 received 4
worker 0 received 9
worker 0 received 14
worker 0 received 18
*/
func test6() {
	var c1, c2 = generator(), generator()
	var worker = createWorker2(0)
	n := 0
	hasValue := false
	for {
		var activeWorker chan<- int //空channel
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n:
			hasValue = false
		}
	}
}

//测试7:
//功能1:解决test6中worker处理太慢造成,很多任务丢失的问题
/*
输出:
worker 0 received 0---
worker 0 received 0  |
worker 0 received 1  |
worker 0 received 1  |
worker 0 received 2  generator()生成速度小于1秒
worker 0 received 2  |
worker 0 received 3  |
worker 0 received 3  |
worker 0 received 4---
worker 0 received 5 //例如:generator()生成的速度1.4 (乱序的原因)
worker 0 received 4 //例如:generator()生成的速度1.5
worker 0 received 5
*/
func test7() {
	var c1, c2 = generator(), generator()
	var worker = createWorker2(0)
	var values []int
	for {
		var activeWorker chan<- int //空channel
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		}
		//打印当前values队列长度
		//fmt.Println("当前values内容:", values)
	}
}

func main() {
	//test6()
	test7()
}
