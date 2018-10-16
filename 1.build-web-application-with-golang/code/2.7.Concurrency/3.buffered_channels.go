package main

import (
	"fmt"
)

//Go允许指定channel的缓冲大小，就是channel可以存储多少元素。
//ch:= make(chan bool, 4)，创建了可以存储4个元素的bool 型channel。
//在这个channel 中，前4个元素可以无阻塞的写入。
//当写入第5个元素时，代码将会阻塞，直到其他goroutine从channel 中读取一些元素，腾出空间。

//ch := make(chan type, value)
//value = 0 时，channel 是无缓冲阻塞读写的，
//当value > 0 时，channel 有缓冲、是非阻塞的，直到写满 value 个元素才阻塞写入。

func test1() {
	//修改为1:ERROR
	//fatal error: all goroutines are asleep - deadlock!
	//修改为2或者3:OK
	c := make(chan int, 2)
	c <- 1 //发东西到channel
	c <- 2

	fmt.Println(<-c) //从channel读取
	fmt.Println(<-c)
}

func main() {
	test1()
}
