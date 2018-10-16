package main

import (
	"fmt"
)

//from C version
func fibnacci_recursive(n int) int {
	if n <= 1 {
		return 1
	} else {
		return fibnacci_recursive(n-1) + fibnacci_recursive(n-2)
	}
}

//from C version
func fibnacci_for(n int) int {
	num := 0
	p1, p2 := 1, 1
	if n < 1 {
		return 1
	}
	for i := 2; i <= n; i++ {
		num = p1 + p2
		p1 = p2
		p2 = num
	}
	return num
}

//c chan int:
//c:是名字
//chan:是说是信道
//int:是说是信道中元素的类型
func fibnacci_producer(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i <= n; i++ {
		c <- x
		x, y = y, x + y
	}
	close(c) //必须close,推荐在生产者关闭
}

//测试最基本的fibnacci功能
func test1() {
	val := 0
	val = fibnacci_recursive(10)
	fmt.Println("val:", val)
	val = fibnacci_for(10)
	fmt.Println("val:", val)
}

//测试Gorutines+channel来实现fibnacci
func test2() {
	c := make(chan int, 20)
	go fibnacci_producer(cap(c), c)

	//第一种写法:
	for {
		if v, ok := <-c; ok {
			fmt.Println(v)
		} else {
			//ok为false说明channel关闭并且没有数据
			fmt.Println("channel is closed and habe not elements")
			break
		}
	}

	////第二种写法:
	//for data := range c {
	//	fmt.Println(data)
	//}
}

func main() {
	//test1()
	fmt.Println("\n===============\n")
	test2()
}
