package main

import (
	"fmt"

	queue "go_study/3.google_deep_go/5.interface/2.queue"
)

//测试API1:
func test1() {
	q := queue.Queue{1}
	q.Push(2) //值做为接受者调用Push
	q.Push(3)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	q.Push("abc")
	fmt.Println(q.Pop())
}

//测试API2和API3:
func test2() {
	q := queue.Queue{1}
	q.Push2(2)
	fmt.Println(q.Pop2())
	fmt.Println(q.Pop2())
	fmt.Println(q.IsEmpty())

	//编译错误
	//q.Push2("abc")

	//运行时错误
	q.Push3("abc")

}

func main() {
	test1()
	fmt.Println("------------------")
	test2()
}
