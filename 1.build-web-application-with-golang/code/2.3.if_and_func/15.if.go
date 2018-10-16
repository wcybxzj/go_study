package main

import "fmt"

//测试1:最普通if
func test1() {
	x := 10
	if x > 10 {
		fmt.Println("ok!")
	} else {
		fmt.Println("not!")
	}
}

func test2() int {
	return 12
}

//测试3:if中有函数执行
func test3() {
	if x := test2(); x > 10 {
		//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
		fmt.Println(x)
		fmt.Println("ok!")
	} else {
		fmt.Println(x)
		fmt.Println("not!")
	}
	//这个地方如果这样调用就编译出错了，因为x是条件里面的变量
	//fmt.Println(x)
}

//测试4:多个if
func test4() {
	integer := 1
	if integer == 3 {
		fmt.Println("equal 3!")
	} else if integer < 3 {
		fmt.Println("less 3!")
	} else {
		fmt.Println("great 3!")
	}
}

func main() {
	test1()
	test3()
	test4()
}
