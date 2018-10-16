package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	num := runtime.NumCPU() //获取cpu核数
	fmt.Println("cpu number is:", num)

	//runtime.GOMAXPROCS(num-1) //制定能使用几个cpu

	for i := 0; i < 10; i++ {
		go func() {
			for {
			}
		}()
	}
	time.Sleep(time.Second * 100)
}
