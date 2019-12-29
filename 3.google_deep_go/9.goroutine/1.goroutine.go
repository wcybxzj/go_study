package main

import (
	"fmt"
	_ "net/http/pprof"
	"runtime"
	"time"
)

//测试1:
//第一种出让协程调度的方法
//IO操作可以让协程交出协程的控制权,所以在4核上启动1000个协程和主协程都有机会去执行
//1毫秒中让1000个写成来并发打印,1毫秒后程序终止
func test1() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				//io操作会让协程进行切换
				fmt.Printf("Hello from  goroutine %d\n", i)
			}
		}(i)
	}
	//和c语言一样必须wait或者pthread_join
	time.Sleep(time.Millisecond)
}

//测试2:
//启动10个协程,4个cpu上测试的,4个cpu都100%,其余6个协程和主协程没机会执行
//协程层面:
//++操作,4个协程之间没机会切换,分别死循环
//主程序层面:
//因为没机会执行所以程序结束不了
func test2() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

//测试3:
//描述:
//第二种出让调度:runtime.Gosched主动出让调度
//其余6个协程和主协程有机会执行
//========================================================
//输出:
//[690 691 762 623 670 581 548 607 525 567]
//========================================================
//潜在的并发问题
//go run -race 1.goroutine.go
//WARNING: DATA RACE
//Read at 0x00c4200b2000 by main goroutine:冲突的读方式
//Previous write at 0x00c4200b2000 by goroutine 6:冲突的写方
//冲突的数据是变量a,冲突的双方分是协程++操作和主协程的Println操作
//解决这个问题需要用后边的channel
//========================================================
//========================================================
func test3() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(ii int) {
			for {
				a[ii]++
				runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

//测试4:
//修改点:
//go协程不传参直接用全局的i
//====================================================
//运行报错:
// panic: runtime error: index out of range
//====================================================
//调试方法:
//go run -race xxxx.go
//说存在DATA RACE
//====================================================
//分析:
//协程的匿名函数使用外部的i,主协程for结束时i=10,但10个子协程引用的还是这个i,立刻数组越界报错
//====================================================
//解决方法:
//使用test3()的写法,对每个i传入到匿名函数里
//====================================================
func test4() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func() {
			for {
				a[i]++
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}

//测试5:
//启动1000个协程，看需要多少线程支撑
func test5() {
	goroutine_num()

	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from  goroutine %d\n", i)
				time.Sleep(time.Second)
			}
		}(i)
	}
	time.Sleep(time.Minute * 10) //10分钟
}

//测试6:
//本机器是4个核心
//启动4个子协程,都是循环看能跑满几个核心?
//直接可以跑满4个核心
func test6() {
	goroutine_num()

	for i := 0; i < 4; i++ {
		go func(i int) {
			for {
				fmt.Printf("Hello from  goroutine %d\n", i)
				time.Sleep(time.Second)
			}
		}(i)
	}
	time.Sleep(time.Minute * 10) //10分钟
}

//测试7:
//100个生产者不断的发任务
//1个消费者每秒只能消费一个任务
//1个协程什么也不做
//看到102个协程被挤压
func test7() {
	goroutine_num()
	c1 := make(chan int)

	for i := 1; i < 100; i++ {
		go func(i int) {
			for {
				c1 <- i
			}
		}(i)
	}

	go func() {
		for {
			i2 := <-c1
			time.Sleep(time.Second)
			println("fired:", i2)
		}
	}()

	//主协程不能退出
	time.Sleep(time.Minute * 10) //10分钟
}

//查询当前协程总数
func goroutine_num() {
	go func() {
		for {
			num := runtime.NumGoroutine()
			fmt.Println("goroutine 总数:", num)
		}
	}()
}

func main() {
	//test1()
	//test2()
	test3()
	//test4()
	//test5()
	//test6()
	//test7()
}
