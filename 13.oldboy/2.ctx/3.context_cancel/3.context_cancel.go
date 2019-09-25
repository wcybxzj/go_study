package __context_cancel

import (
	"context"
	"fmt"
	time2 "go/src/time"
	"time"
)

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("2222222222222222222222222222222")
				fmt.Println("gen exited ")
				return // returning not to leak the goroutine
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

//此示例演示使用一个可取消的上下文，以防止 goroutine 泄漏。示例函数结束时，defer 调用 cancel 方法，gen goroutine 将返回而不泄漏。
//通过cancel来取消协程,避免协程泄露
//执行顺序:
// main协程test()结束触发defer cancel()
// 子协程发现Done发现cancel()子协程结束

/*
输出:
1
2
3
4
5
1111111111111111111111111111111
2222222222222222222222222222222
gen exited
*/
func test() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	fmt.Println("1111111111111111111111111111111")
}

func main() {
	test()
	time.Sleep(time2.Microsecond)
}
