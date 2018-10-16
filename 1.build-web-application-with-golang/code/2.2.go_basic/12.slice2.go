package main

import "fmt"

//len 获取slice的长度
//cap 获取slice的最大容量
func test1() {
	arr1 := [10]int{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var s1 []int
	s1 = arr1[:3]
	printSlice(s1)                  //len=2 cap=8 slice=[22 33]
	fmt.Printf("len:%d\n", len(s1)) //len:3
	fmt.Printf("cap:%d\n", cap(s1)) //cap:10
}

//append 向slice里面追加一个或者多个元素，然后返回一个和slice一样类型的slice
//copy 函数copy从源slice的src中复制元素到目标dst，并且返回复制的元素的个数
func test2() {
	var numbers []int
	/* 向切片添加一个元素 */
	printSlice(numbers) //len=0 cap=0 slice=[]
	numbers = append(numbers, 123)
	printSlice(numbers) //len=1 cap=1 slice=[123]

	/* 同时添加多个元素 */
	//输出:len=4 cap=4 slice=[123 2 3 4]
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	//输出:len=4 cap=8 slice=[0 0 0 0]
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)
	printSlice(numbers1)

	/* 拷贝 numbers 的内容到 numbers1 */
	//len=4 cap=8 slice=[123 2 3 4]
	copy(numbers1, numbers)
	printSlice(numbers1)

}

//https://github.com/astaxie/build-web-application-with-golang/blob/master/zh/02.2.md
//图中jie解释了 cap为什么是8和5
func test3() {
	arr1 := [10]int{10, 11, 22, 33, 44, 55, 66, 77, 88, 99}
	var s1 []int
	var s2 []int

	//len=2 cap=8 slice=[22 33]
	s1 = arr1[2:4]
	printSlice(s1)

	//len=2 cap=5 slice=[22 33]
	s2 = arr1[2:4:7]
	printSlice(s2)
}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	test1()
	test2()
	test3()
}
