package main

import (
	"fmt"
)

func test1() {
	//空接口像c语言 void *可以指向任务类型
	//a为 空借口
	var a interface{}
	var i int = 5
	s := "hello world"

	a = i
	fmt.Printf("%d\n", a)

	a = s
	fmt.Printf("%s\n", a)
}

func main() {
	test1()
}
