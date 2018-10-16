package main

import "fmt"

func test1() {
	v1, v2, v3 := 11, 22, 33
	fmt.Print(v1, v2, v3)
	fmt.Print("\n")
}

func test2() {
	//_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。
	//在这个例子中，我们将值35赋予b，并同时丢弃34
	_, b := 34, 35
	fmt.Print(b)
	fmt.Print("\n")
}

func main() {
	test1()
	test2()
}
