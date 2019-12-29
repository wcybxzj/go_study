package main

import (
	"context"
	"fmt"
	"time"
)

/*
输出:
context deadline exceeded

50ms WithDeadline将会触发
*/

func test1()  {
	d := time.Now().Add(time.Second)

	ctx, cancel := context.WithDeadline(context.Background(), d)

	// Even though ctx will be expired, it is good practice to call its
	// cancelation function in any case. Failure to do so may keep the
	// context and its parent alive longer than necessary.
	defer cancel()

	for {
		select {
		case <-time.After(100 * time.Millisecond):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println("ctx.Done")
			fmt.Println(ctx.Err())
			return
		}
	}
}

func main() {
	 test1()
}
