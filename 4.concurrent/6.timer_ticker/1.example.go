package main

import (
	"fmt"
	"time"
)

/*
结果:死锁

start 1 th worker
start 2 th worker
start 3 th worker
1
2
3
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
	/root/www/go_www/src/go_study/4.concurrent/6.timer_ticker/1.example.go:50 +0xc8

goroutine 19 [chan receive]:
main.DoTickerWork.func1(0xc420092000, 0xc420088140)
	/root/www/go_www/src/go_study/4.concurrent/6.timer_ticker/1.example.go:35 +0x78
created by main.DoTickerWork
	/root/www/go_www/src/go_study/4.concurrent/6.timer_ticker/1.example.go:31 +0x5c
*/
func DoTickerWork(res chan interface{}, timeout <-chan time.Time) {
	t := time.NewTicker(3 * time.Second)
	go func() {
		defer close(res)
		i := 1
		for {
			<-t.C
			fmt.Printf("start %d th worker\n", i)
			res <- i
			i++
		}
	}()
	<-timeout
	t.Stop()
	return
}

func main() {
	res := make(chan interface{}, 10000)
	timeout := time.After(10 * time.Second)
	DoTickerWork(res, timeout)
	for v := range res {
		fmt.Println(v)
	}
}
