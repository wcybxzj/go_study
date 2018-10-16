package main

import (
	"fmt"
	"runtime"
)

//runtime.Gosched()表示让CPU把时间片让给别人,下次某个时候继续恢复执行该goroutine。
func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched() //相当于c语言的多线程中的sched_yield()
		fmt.Println(s)
		//fmt.Printf("Hello from  goroutine %s\n", s)
	}
}

//多个goroutine运行在同一个进程里面，共享内存数据，
//设计上要遵循:不要通过共享来通信，而通过通信来共享。
func test1() {
	go say("world") //开一个新的Goroutins执行
	say("hello")    //当前Goroutines执行

}

func test2() {
	fmt.Println("机器CPU数量:", runtime.NumCPU())
	//runtime.GOMAXPROCS(4)
	//fmt.Println()
	//fmt.Printf("%d\n", runtime.GOMAXPROCS)
}

func main() {
	//test1()
	test2()
}
