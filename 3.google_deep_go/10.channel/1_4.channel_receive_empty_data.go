package main

import (
	"fmt"
	"time"
)

//现象:消费者读出大量空值
//原因:需要生产者close channel,消费者用最简单的for来读channel,无法得知channel已经关闭
//解决办法:消费者用for+range或者for+if

//父生产,子消费:消费者读取出大量空值的情况
//实现读出大量空值:1_1.channel.go test8()使用worker7()
//解决办法:1_1.channel.go test8()使用worker8_1()或者worker8_2()
func test1() {

}

func producer2(a ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range a {
			out <- v
		}
		close(out)
	}()
	return out
}

//子生产,父消费:消费者读取出大量空值的情况
//输出:3,2,6,7,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
func test2() {
	out := producer2(3, 2, 6, 7, 4)

	for {
		num := <-out
		fmt.Print(num)
		fmt.Print(",")
		time.Sleep(time.Second / 4)
	}
}

func main() {
	test2()
}
