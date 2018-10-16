package main

import (
	"fmt"
	"time"
)

//测试1:(失败)
//结果:死锁
//原因:只有人在写channel,没人在收数据
func test1() {
	//var c chan int //c == nil
	c := make(chan int)
	c <- 1 //死锁原因发数据没有其他协程来收数据
	c <- 2
	n := <-c
	fmt.Println(n)
}

//测试2:(失败)
//结果:只打印出1
//原因:子协程还没打印出内容,主协程就结束,线程也结束了
func test2() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
}

//测试3:(成功)
//结果:打印出1和2
func test3() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

//测试4:(成功)
//channel做参数进行传递
//协程抽取出一个普通函数
//结果:打印出1和2
func worker4(c chan int) {
	for {
		n := <-c
		fmt.Println(n)
	}
}

func test4() {
	c := make(chan int)
	go worker4(c)
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
}

//测试5:(成功)
//channel做为数组
//输出:
/*
worker 0 received a
worker 0 received A
worker 1 received b
worker 1 received B
worker 2 received c
worker 2 received C
worker 3 received d
worker 3 received D
worker 4 received e
worker 4 received E
*/
func worker5(id int, c chan int) {
	for {
		fmt.Printf("worker %d received %c\n", id, <-c)
	}
}

func test5() {
	var channels [5]chan int //声明
	for i := 0; i < 5; i++ {
		channels[i] = make(chan int) //创建
		go worker5(i, channels[i])
	}

	for i := 0; i < 5; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 5; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

//测试6:(成功)
//channel作为返回值:
//chan<-int意思是:只能往channel里写
func createWorker6(id int) chan<- int {
	c := make(chan int)
	go func() {
		for {
			fmt.Printf("worker %d received %c\n", id, <-c)
		}
	}()
	return c
}

func test6() {
	var channels [5]chan<- int //声明一个只写的channel
	for i := 0; i < 5; i++ {
		channels[i] = createWorker6(i)
	}

	for i := 0; i < 5; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 5; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)
}

//测试7:(成功)
//bufferedChannel:
//用处1:用来解决test1中的问题
//用处2:可以提高性能
//增加channel缓冲,缓冲满前即使没人来读,也可以往里写而不阻塞
func worker7(id int, c chan int) {
	for {
		//fmt.Printf("Worker %d received %c \n", id, <-c)

		cNum := <-c
		fmt.Printf("Worker %d received %c \n", id, cNum)
	}
}

func test7() {
	c := make(chan int, 3)
	go worker7(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

//测试8:
//channelClose()

//test8()使用worker7(失败)
// 没有识别发送方是否channelClose ,worker会接受大量空值在正常数据后边
/*输出:
Worker 0 received a
Worker 0 received b
Worker 0 received c
Worker 0 received d
Worker 0 received
Worker 0 received
Worker 0 received
....................
*/

//test8()使用worker8_1或者worker8_2 (成功)
// 能够识别发送方已经close的情况
/*输出
Worker 0 received a
Worker 0 received b
Worker 0 received c
Worker 0 received d
*/

//方式1:
//if判断发送方是否关闭channel,接收方关闭channel时候break
func worker8_1(id int, c chan int) {
	for {
		n, ok := <-c
		if !ok {
			break
		}
		fmt.Printf("Worker %d received %c \n", id, n)
	}
}

//方式2:
//range自动判断发送方是否关闭channel,接收方关闭channel时候自动退出for
func worker8_2(id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %c \n", id, n)
	}
}

func test8() {
	c := make(chan int)
	go worker7(0, c) //fail
	//go worker8_1(0, c) //ok
	//go worker8_2(0, c) //ok
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
	time.Sleep(time.Millisecond)
}

//测试9:(正常)
//不使用channelClose
//发送方不channelClose,接受方使用range来接受channel数据
//最后是worker8_2,range都读不出数据阻塞,主协程运行完退出,线程退出,所有协程也退出
func test9() {
	c := make(chan int)
	go worker7(0, c) //ok
	//go worker8_1(0, c) //ok
	//go worker8_2(0, c) //ok
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	time.Sleep(time.Millisecond)
}

func main() {
	//test1()
	//test2()
	//test3()
	//test4()
	//test5()
	//test6()
	//test7()
	//test8()
	test9()
}
