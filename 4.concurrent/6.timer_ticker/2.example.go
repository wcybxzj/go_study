package main

import (
	"fmt"
	"time"
)

func DoTickerWork2(res chan interface{}, timeout <-chan time.Time) {
	t := time.NewTicker(3 * time.Second)
	done := make(chan bool, 1)
	go func() {
		defer close(res)
		i := 1
		for {
			select {
			case <-t.C:
				fmt.Printf("start %d th worker\n", i)
				res <- i
				i++
			case <-timeout:
				close(done)
				return
			}
		}
	}()
	<-done
	return
}

func main() {
	res := make(chan interface{}, 10000)
	timeout := time.After(10 * time.Second)
	DoTickerWork2(res, timeout)
	for v := range res {
		fmt.Println(v)
	}
}
