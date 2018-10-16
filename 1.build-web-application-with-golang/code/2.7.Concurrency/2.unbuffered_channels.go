package main

import (
	"fmt"
)

//func test1() {
//	c1 := make(chan int)
//	c2 := make(chan string)
//	c3 := make(chan interface{})
//}

//channel:
//1.channel是int类型
//2.channel必须用make创建
//3.channel分析
//默认情况，channel接收和发送数据都是阻塞的.
//除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。
//所谓阻塞，
//如果读,（value := <-ch）它将会被阻塞，直到有数据接收。
//如果写,任何发送（ch<-5）将会被阻塞，直到数据被读出。
//无缓冲channel是在多个goroutine之间同步很棒的工具。
func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total //写入channel
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)

	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}
