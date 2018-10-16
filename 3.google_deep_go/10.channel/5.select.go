package main

import (
	"fmt"
	"math/rand"
	"time"
)

//测试0:
//情况:向nil channel写
//结果:死锁
func test0() {
	var activeWorker chan<- int //空channel
	n := 123
	select {
	case activeWorker <- n:
		fmt.Println(123)
	}
}

//测试1:
//情况:从nil channel读
//结果:死锁
func test1() {
	//两个channel同时进行接收,谁来的快收.
	var c1, c2 chan int //声明后都是nil
	select {
	case n := <-c1:
		fmt.Println("received from c1:", n)
	case n := <-c2:
		fmt.Println("received from c2:", n)
	}
}

//测试2:select+default实现非阻塞使用channel
func test2() {
	var c1, c2 chan int
	for {
		select {
		case n := <-c1:
			fmt.Println("received from c1:", n)
		case n := <-c2:
			fmt.Println("received from c2:", n)
		default:
			fmt.Println("no value received ")
		}
	}
}

//1秒产生1个数字
func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1000)) *
				time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker(id int, c chan int) {
	for n := range c {
		fmt.Printf("worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

//测试3:(成功)
//两个channel 任何一个来数据就会读出来
func test3() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)

	for {
		n := 0
		select {
		case n = <-c1://读
			w <- n //写
		case n = <-c2:
			w <- n
		}
	}
}

//测试4:(问题版)
//让select既可以从c1,c2收数据,也可以向c发数据
//问题原因:n还没读出来就拼命往w里写0
/*
输出:
worker 0 received 0
worker 0 received 0
worker 0 received 0
worker 0 received 0
worker 0 received 0
..................
*/
func test4() {
	var c1, c2 = generator(), generator()
	w := createWorker(0)

	n := 0
	for {
		select {
		case n = <-c1: //因为可能还没生成数据,所以没机会执行
		case n = <-c2: //因为可能还没生成数据,所以没机会执行
		case w <- n: //一直把0写入到这个channel
		}
	}
}

//测试5:(成功)
//1.解决test4()的问题:
//从c1和c2中收到数据前,activeWorker是nil channel,无法写入
//从c1和c2中收到数据后,activeWorker成为普通channel,才可以写入
func test5() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)
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

func main() {
	//test0()
	test1()
	//test2()
	//test3()
	//test4()
	//test5()
}
