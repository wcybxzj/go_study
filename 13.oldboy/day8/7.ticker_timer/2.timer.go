package main

import (
	"fmt"
	"runtime"
	"time"
)

//内存泄露在linux没实现,在视频里的windows也没实现
//time.After:启动一个协程,往channel里写内容,返回这个channel
//不推荐使用time.After
func test1() {
	num := runtime.NumCPU()
	runtime.GOMAXPROCS(num - 1)

	fmt.Println("cpu number is:", num)

	for i := 0; i < 10; i++ {
		go func() {
			for {
				select {
				case <-time.After(time.Second):
					//fmt.Println("alert")
				}
			}
		}()
	}
	time.Sleep(time.Second * 100)
}

//推荐使用 NewTimer()+stop来使用
func test2() {
	for i := 0; i < 1; i++ {
		go func() {
			for {
				t := time.NewTimer(time.Second)
				select {
				case <-t.C:
					fmt.Println("timeout")
				}
				t.Stop()
			}
		}()
	}

	time.Sleep(time.Second * 100)
}

func main() {
	//test1()
	test2()
}
