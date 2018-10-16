package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
	LOOP:
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second)
			fmt.Println(i)
			if i == 3 {
				break LOOP
			}
		}

		fmt.Println("123")
	}()

	time.Sleep(time.Second * 100)
}
