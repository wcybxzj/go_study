package main

import (
	"fmt"
)

//go语言的指针:
//go指针不能运算,
//不像c语言指针回去ptr++,移动的多少要看指针的类型是什么

//go参数传递:
//go只有值传递,意味着单独用值传递数据是要拷贝的
//所以要使用指针来避免数据复制,而是去复制地址，
//C也是这样啊

//测试1:完全和C一样
func test1() {
	a := 2
	var pa *int = &a
	*pa = 123
	fmt.Println(a)//123

	b := 456
	*pa = b
	fmt.Println(a)//456
}

//使用指针swap
func swap(a, b *int) {
	*b, *a = *a, *b
}

//不使用指针swap
func swap2(a, b int) (int, int) {
	return b, a
}

//测试2:和C完全一样
func test2() {
	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)//4, 3

	a, b = 5, 6
	a, b = swap2(a, b)
	fmt.Println(a, b)//6, 5
}

func main() {
	test1()
	fmt.Println("==================")
	test2()
}
