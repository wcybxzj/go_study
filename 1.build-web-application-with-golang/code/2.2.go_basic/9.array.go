package main

import "fmt"

func test1() {
	var arr [10]int
	arr[0] = 11
	arr[1] = 22
	arr[2] = 33
	arr[3] = 44
	arr[4] = 55

	fmt.Printf("%d\n", arr[0])
	fmt.Printf("%d\n", arr[1])
	fmt.Printf("%d\n", arr[2])
	fmt.Printf("%d\n", arr[3])
	fmt.Printf("%d\n", arr[4])
	fmt.Printf("%d\n", arr[9]) //返回未赋值的最后一个元素，默认返回0
}

func test2() {
	a := [3]int{1, 2, 3}   // 声明了一个长度为3的int数组
	b := [10]int{1, 2, 3}  // 声明了一个长度为10的int数组，其中前三个元素初始化为1、2、3，其它默认为0
	c := [...]int{4, 5, 6} // 可以省略长度而采用`...`的方式，Go会自动根据元素个数来计算长度

	fmt.Printf("%d, %d, %d\n", a[0], a[1], a[2])
	fmt.Printf("%d, %d, %d\n", b[0], b[1], b[2])
	fmt.Printf("%d, %d, %d\n", c[0], c[1], c[2])

}

//数组不是引用类型
//输出:
//[123 22 33]
//[11 22 33]
func test3() {
	a1 := [3]int{11, 22, 33}
	a2 := a1
	a1[0] = 123
	fmt.Printf("%v\n", a1)
	fmt.Printf("%v\n", a2)
}

func main() {
	test1()
	fmt.Printf("===============\n")
	test2()
	fmt.Printf("===============\n")
	test3()
}
