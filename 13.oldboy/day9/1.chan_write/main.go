package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	ch = make(chan int, 1)

	go func() {
		var i int
		for {
			select {
			case ch <- i:
				fmt.Println("send ok i:", i)
			default:
				fmt.Println("channel is full i", i)
				time.Sleep(time.Second)
			}
			i++
		}
	}()

	for {
		v := <-ch
		fmt.Println(v)
	}
}
