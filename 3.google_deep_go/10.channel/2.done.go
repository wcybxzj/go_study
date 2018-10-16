package main

import (
	"fmt"
)

func doWork(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork(id, w.in, w.done)
	return w
}

//测试1:(失败)
//解决:
//在1.channel.go:test5()中,主协程用sleep来灯子协程执行完毕
//通过增加一个done channel,让子协程去通知父协程自己工作已经完成
//新的问题:
//程序成为串行而不是并行了,看worker序号
/*
输出:(和1.channel.go test5()比较)
worker 0 received a
worker 1 received b
worker 2 received c
worker 3 received d
worker 4 received e
worker 5 received f
worker 6 received g
worker 7 received h
worker 8 received i
worker 9 received j
worker 0 received A
worker 1 received B
worker 2 received C
worker 3 received D
worker 4 received E
worker 5 received F
worker 6 received G
worker 7 received H
worker 8 received I
worker 9 received J
*/
func test1() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i //发1个任务
		<-workers[i].done        //等1个任务执行完
	}

	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
		<-workers[i].done
	}
}

//测试2:(失败)
//想要优化test1()实现并行,但是死锁
/*
输出:
worker 0 received a
worker 1 received b
worker 2 received c
worker 3 received d
worker 4 received e
worker 5 received f
worker 6 received g
worker 7 received h
worker 8 received i
worker 9 received j
fatal error: all goroutines are asleep - deadlock!

情况描述:
	1.worker的处理时乱序的说明是并发
	2.所有小写字母都打印出来了
分析:
	父协程向channel发送小写字母
	子协程从channel读取了小写字母
	子协程向done发送true(没人读死锁)
	父协程向channel发送大写字母(没人读死锁)
*/
func test2() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	//发送所有任务
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	//等所有子协程执行完毕
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

func doWork3(id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("worker %d received %c\n", id, n)
		//协程中的协程
		go func() {
			done <- true
		}()
	}
}

func createWorker3(id int) worker {
	w := worker{
		in:   make(chan int),
		done: make(chan bool),
	}
	go doWork3(id, w.in, w.done)
	return w
}

//测试3:(成功)
//解决test2()死锁的第一种方法:
//让子协程里再开1个协程,单独去往done去写
//结果:worker是乱序的说明是并发(仔细往后看)
/*
输出:
worker 0 received a
worker 0 received A
worker 1 received b
worker 1 received B
worker 2 received c //开发出现乱序
worker 3 received d
worker 4 received e
worker 5 received f
worker 6 received g
worker 7 received h
worker 8 received i
worker 9 received j
worker 2 received C
worker 3 received D
worker 4 received E
worker 5 received F
worker 6 received G
worker 7 received H
worker 8 received I
worker 9 received J

分析:
	父协程向channel发送小写字母
	子协程从channel读取了小写字母
	子协程单独开一个协程向done发送true(现在还没人读,所以阻塞)
	父协程向channel发送大写字母
	子协程从channel读取了大写字母
	子协程单独开一个协程向done发送true(现在还没人读,所以阻塞)
	父协程从done读取所有的信息
*/
func test3() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker3(i)
	}

	//发送所有任务
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	//等所有子协程执行完毕
	for _, worker := range workers {
		<-worker.done
		<-worker.done
	}
}

//测试4:(成功)
//解决test2()死锁的第二种方法:
/*
worker 0 received a
worker 1 received b
worker 2 received c
worker 3 received d
worker 4 received e
worker 5 received f
worker 6 received g
worker 7 received h
worker 8 received i
worker 9 received j
worker 0 received A
worker 1 received B
worker 2 received C
worker 3 received D
worker 4 received E
worker 5 received F
worker 6 received G
worker 7 received H
worker 8 received I
worker 9 received J
*/
func test4() {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	//发送所有任务
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for _, worker := range workers {
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	for _, worker := range workers {
		<-worker.done
	}
}

func main() {
	//test1()
	//test2()
	test3()

	//test4()
}
