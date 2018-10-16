package main

import (
	"fmt"
	"go_study/3.google_deep_go/4_3.extend_type/2.queue"
)

func main() {
	q := queue.Queue{1}
	q.Push(2) //值做为接受者调用Push
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
