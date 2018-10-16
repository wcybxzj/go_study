package main

import "fmt"

//break操作是跳出当前循环，continue是跳过本次循环。
//当嵌套过深的时候，break可以配合标签使用，即跳转至标签所指定的位置

func test1() {
	for index := 10; index > 0; index-- {
		if index == 5 {
			break
		}
		fmt.Println(index)
	}
}

func test2() {
	for index := 10; index > 0; index-- {
		if index == 5 {
			continue
		}
		fmt.Println(index)
	}
}

func main() {
	test1()
	fmt.Println("\n")
	test2()
	fmt.Println("\n")
}
