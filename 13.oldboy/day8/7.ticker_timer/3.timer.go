package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	go func() {
		var i int
		for {
			ch1 <- i
			time.Sleep(time.Second)

			ch2 <- i * i
			time.Sleep(time.Second)
			i++
		}
	}()

	for {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		case <-time.After(time.Second):
			fmt.Println("get data timeout")
			time.Sleep(time.Second)
		}
	}
}
