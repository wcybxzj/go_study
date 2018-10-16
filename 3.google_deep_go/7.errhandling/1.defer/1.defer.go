package main

import (
	"bufio"
	"fmt"
	"go_study/google_deep_go/functional/fib"
	"os"
)

//输出:2,1
func tryDefer() {
	defer fmt.Println(1)
	fmt.Println(2)
}

//输出:3,2,1
//说明:defer是栈结构
func tryDefer2() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

//输出:3,2,1
func tryDefer3() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	return
	fmt.Println(4)
}

//输出:3,2,1
func tryDefer4() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	panic("error occurred")
	fmt.Println(4)
}

/*
输出:
3
2
1
0
panic: is is bigger than 3
*/
//说明参数在defer语句时计算,而不是在defer语句执行时计算
func tryDefer5() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
		if i == 3 {
			panic("is is bigger than 3")
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//直接写文件比较慢,使用bufio
	writer := bufio.NewWriter(file)
	defer writer.Flush() //刷新缓冲到文件

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	//tryDefer()
	//tryDefer2()
	//tryDefer3()
	//tryDefer4()
	tryDefer5()

	//writeFile("123.txt")

}
