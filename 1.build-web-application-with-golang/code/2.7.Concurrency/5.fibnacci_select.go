package main

import (
	"fmt"
)

func fibnacci1(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit: //读channel quit
			return
		default:
			//fmt.Println("主阻塞在这里会疯狂打印")
		}
	}
}

func test1()  {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("子读", <-c) //从c读取
		}
		quit <- 0 //把0写入quit
	}()

	fibnacci1(c, quit)
}


//多个channel时,使用select 选择channel
//select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，
//当多个channel都准备好的时候，select是随机的选择一个执行的
//当多个channel中有1个chanenl准备好的时候,select会用那个channel
func main() {
	test1()
}
