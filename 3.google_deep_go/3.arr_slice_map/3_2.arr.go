package main

import "fmt"

//知识点:
//知识点1:[10]int 和[20]int是不同的类型
//知识点2:func f(arr [10]int) 在使用这个函数的时候,整个数组会拷贝到函数中
//知识点3:go一般不直接使用数组(测试1),也不会去用数组指针做参数(参数2),一般用切片

//测试1:数组做为参数
//GO数组是值类型,传入后就会进行整个数组内容的拷贝,非常危险
//C数组做参数是引用传递,也就是说传的是地址,不会进行整个数组的数据拷贝
/*
[root@localhost 3.google_deep_go]# go run 3_2.arr.go
[0 0 0 0 0] [1 3 5] [2 4 6 8 10]
index:0, value:100
index:1, value:0
index:2, value:0
index:3, value:0
index:4, value:0
index:0, value:100
index:1, value:4
index:2, value:6
index:3, value:8
index:4, value:10
[0 0 0 0 0] [1 3 5] [2 4 6 8 10]
*/
func test1() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3)

	printArray(arr1)
	//printArray(arr2) //报错数组大小不符合函数参数的定义
	printArray(arr3)

	fmt.Println(arr1, arr2, arr3)
}

func printArray(arr [5]int) {
	arr[0] = 100 //说明数组是值传递,原来的arr值并不会改变
	for index, value := range arr {
		fmt.Printf("index:%d, value:%d\n", index, value)
	}
}

//测试2:数组指针做为参数
/*
go run 3_2.arr.go
[0 0 0 0 0] [1 3 5] [2 4 6 8 10]
index:0, value:100
index:1, value:0
index:2, value:0
index:3, value:0
index:4, value:0
index:0, value:100
index:1, value:4
index:2, value:6
index:3, value:8
index:4, value:10
[100 0 0 0 0] [1 3 5] [100 4 6 8 10]
*/
func test2() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}

	fmt.Println(arr1, arr2, arr3)

	//GO不像C语言,
	//C语言中数组名是数组首元素的地址,可以用数组名也可以用&
	//GO要表示地址只能用&
	printArray2(&arr1)
	//printArray(&arr2) //报错数组大小不符合函数参数的定义
	printArray2(&arr3)

	fmt.Println(arr1, arr2, arr3)
}

func printArray2(arr *[5]int) {
	arr[0] = 100 //和printArray写法完全一样,只是由于参数是指针就可以修改arr
	for index, value := range arr {
		fmt.Printf("index:%d, value:%d\n", index, value)
	}
}

func main() {
	//test1()
	test2()
}
