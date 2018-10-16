package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

const num = 1

func do(i int, c chan int) {
	for {
		fmt.Println("do:")
		time.Sleep(time.Second)
	}
	c <- 123
}

func workerController(i int) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	c := make(chan int, 1)
	go do(i, c)
	select {
	case <-ctx.Done():
		fmt.Println("job work timeout! ")
		//cancel()
	case res := <-c:
		fmt.Printf("job work ok res:%v", res)
	}
}

func goroutine_num() {
	go func() {
		for {
			num := runtime.NumGoroutine()
			fmt.Println("goroutine 总数:", num)
			time.Sleep(time.Second)
		}
	}()
}

//失败
//超时后子协程还在跑
func main() {
	go goroutine_num()

	for i := 0; i < num; i++ {
		go workerController(i)
	}

	time.Sleep(time.Second * 30000)
}
