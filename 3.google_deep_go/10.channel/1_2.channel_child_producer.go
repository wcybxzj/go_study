package main

import (
	"fmt"
	"time"
)

func ArraySourceNotClose(a ...int) chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
	}()

	return out
}

func ArraySourceClose(a ...int) chan int {
	out := make(chan int)

	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()

	return out
}

//测试0:
//---------------------------------------------------
//ArraySourceNotClose():---失败
//报错:
//3,2,6,7,4,fatal error: all goroutines are asleep - deadlock!
//goroutine 1 [chan receive]
//---------------------------------------------------
//ArraySourceClose():---失败
//错误:无法识别对方已经关闭channel,读到一堆0
//3,2,6,7,4,0,0,0,0,0,....
func test0() {
	p := ArraySourceNotClose(3, 2, 6, 7, 4)
	//p := ArraySourceClose(3, 2, 6, 7, 4)

	for {
		num := <-p
		fmt.Print(num)
		fmt.Print(",")
		time.Sleep(time.Second / 4)
	}
}

//测试1:for+if
//---------------------------------------------------
//ArraySourceNotClose():---失败
//报错:
//fatal error: all goroutines are asleep - deadlock!
//goroutine 1 [chan receive]
//---------------------------------------------------
// ArraySourceClose():---成功
func test1() {
	//p := ArraySourceNotClose(3, 2, 6, 7, 4)
	p := ArraySourceClose(3, 2, 6, 7, 4)

	for {
		if num, ok := <-p; ok {
			fmt.Print(num)
			fmt.Print(",")
		} else {
			//不break,会死循环
			break
		}
		time.Sleep(time.Second / 4)
	}
}

//测试2:for+range更简练,效果同for+if
//---------------------------------------------------
//ArraySourceNotClose():---失败
//报错:
//fatal error: all goroutines are asleep - deadlock!
//goroutine 1 [chan receive]
//---------------------------------------------------
// ArraySourceClose():---成功

func test2() {
	//p := ArraySourceNotClose(3, 2, 6, 7, 4)
	p := ArraySourceClose(3, 2, 6, 7, 4)

	for v := range p {
		fmt.Println(v)
	}
}

//子为生产者,父为消费者
//结论1:
//子做生产者发送完数据确定一定要close channel,否则无论消费者如何弄都会死锁
//结论2:
//子做生产者close channel,消费者要用for+if或者for+range来读取channel,才能识别出close channel,否则读出大量空数据

func main() {
	test0()
	//test1()
	//test2()
}
