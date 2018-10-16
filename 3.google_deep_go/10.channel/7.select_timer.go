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

//测试8:
//定时器1:
//给程序设置定时器,10秒后程序终止
//===================================================================
//===================================================================
//定时器2:
//如果两次generator生成任务的间隔超过800ms测打印timeout
//严格来说是两次从for的channel中获取数据的时间间隔
//===================================================================
//===================================================================
//定时器3:
//每隔一定时间打印队列内容
//===================================================================
//如果定时器3设置为每1秒打印队列内容,
//定时器2就很难触发因为tick每1秒触发了1次,很难出现800ms没有任何队列有消息
/*
输出:
当前values队列内容: [0 1 1]
worker 0 received 0
当前values队列内容: [1 1 2 2 3]
worker 0 received 0
当前values队列内容: [1 2 2 3 3]
worker 0 received 1
当前values队列内容: [2 2 3 3 4 5 4 5 6 6]
worker 0 received 1
当前values队列内容: [2 3 3 4 5 4 5 6 6 7]
worker 0 received 2
当前values队列内容: [3 3 4 5 4 5 6 6 7 7 8]
worker 0 received 2
当前values队列内容: [3 4 5 4 5 6 6 7 7 8 9 8]
worker 0 received 3
当前values队列内容: [4 5 4 5 6 6 7 7 8 9 8 10 9]
worker 0 received 3
当前values队列内容: [5 4 5 6 6 7 7 8 9 8 10 9 11 10]
worker 0 received 4
after 10 seconds finish
*/
//===================================================================
//如果定时器3设置为每2秒打印队列内容,定时器2就很容易触发
/*
输出:
worker 0 received 0
当前values队列内容: [1 1 2 2 3]
worker 0 received 0
worker 0 received 1
当前values队列内容: [2 2 3 3 4 5 4 5 6 6]
worker 0 received 1
worker 0 received 2
当前values队列内容: [3 3 4 5 4 5 6 6 7 7 8]
worker 0 received 2
timeout
worker 0 received 3
当前values队列内容: [4 5 4 5 6 6 7 7 8 9 8 10 9]
worker 0 received 3
timeout
worker 0 received 4
after 10 seconds finish
*/
func test8() {
	var c1, c2 = generator(), generator()
	var worker = createWorker2(0)
	var values []int

	//定时器1:
	//10秒后返回一个只读channel
	tm := time.After(10 * time.Second)

	//定时器3:
	//tick := time.Tick(time.Second * 1) //1秒
	tick := time.Tick(time.Second * 2) //2秒

	for {
		var activeWorker chan<- int //nil channel
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
		case <-tick: //定时器3
			fmt.Println("当前values队列内容:", values)
		case <-time.After(800 * time.Millisecond): //定时器2
			fmt.Println("timeout")
		case <-tm: //定时器1
			fmt.Println("after 10 seconds finish")
			return
		}
	}
}

func main() {
	test8()
}
