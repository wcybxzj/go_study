package main

import (
	"fmt"
	"time"
)

var out chan int

func producer1(a ...int) {
	go func() {
		for _, v := range a {
			out <- v
		}
	}()
}

func consumer2()  {
	for {
		num := <-out
		fmt.Print(num)
		fmt.Print(",")
		time.Sleep(time.Second / 4)
	}
}

/*
test1:失败,死锁
主协程消费,子协程生产
主协程一直在用for读,最后没数据了

输出:
3,2,6,7,4,fatal error: all goroutines are asleep - deadlock!
goroutine 1 [chan receive]:

解决方法:1_2.channel.go 子协程做生产者必须close channel
*/
func test1()  {
	out = make(chan int)
	go producer1(3, 2, 6, 7, 4)
	for {
		num := <-out
		fmt.Print(num)
		fmt.Print(",")
		time.Sleep(time.Second / 4)
	}
}

/*
test2:成功
子协程消费,主协程生产
主协程生产完退出,整个进程结束
*/
func test2()  {
	out = make(chan int)

	go consumer2()

	a :=[]int{3, 2, 6, 7, 4}
	for _, v := range a {
		out <- v
	}
}

func main() {
	test1()
	//test2()
}
