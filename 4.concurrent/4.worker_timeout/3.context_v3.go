package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

const num = 1

func do(ctx context.Context, i int, c chan int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Printf("Proc-%d \n", i)
			time.Sleep(time.Second * 1)
		}
	}
}

func workerController(i int) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	c := make(chan int, 1)
	go do(ctx, i, c)

	time.Sleep(time.Second * 2)
	cancel()

	time.Sleep(time.Second)
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

//成功,但是不满足需求TODO
//https://zhuanlan.zhihu.com/p/26695984
func main() {
	go goroutine_num()

	for i := 0; i < num; i++ {
		go workerController(i)
	}

	time.Sleep(time.Second * 30000)
}
