package main

import "fmt"

//观察cap的扩容
func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d\n", s, len(s), cap(s))
}

func printSlice2(s []int) {
	fmt.Printf("数据%v len=%d, cap=%d\n", s, len(s), cap(s))
}

//测试1:append()
/*
输出:
len=0, cap=0
len=1, cap=1
len=2, cap=2
len=3, cap=4
len=4, cap=4
len=5, cap=8
len=6, cap=8
len=7, cap=8
len=8, cap=8
len=9, cap=16
len=10, cap=16
len=11, cap=16
------省略-------
*/
func test1() {
	//zero value for slice in nil
	//s=nil
	var s []int

	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}

	fmt.Println(s)
}

//测试2:另外两种创建切片的方法
//数据[2 4 6 8] len=4, cap=4
//数据[0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] len=16, cap=16
//数据[0 0 0 0 0 0 0 0 0 0] len=10, cap=32
func test2() {
	s1 := []int{2, 4, 6, 8}
	printSlice2(s1)

	//可以设置生成切片时cap的大小,避免test1()中数据增加后切片要重新分配空间
	//make(切片类型, len, cap)
	s2 := make([]int, 16)
	printSlice2(s2)

	s3 := make([]int, 10, 32)
	printSlice2(s3)
}

//copy 和 del
//输出
//数据[2 4 6 8 0 0 0 0 0 0 0 0 0 0 0 0] len=16, cap=16
//数据[2 4 6 0 0 0 0 0 0 0 0 0 0 0 0] len=15, cap=16
func test3() {
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)

	fmt.Println("copy():")
	copy(s2, s1)
	printSlice2(s2)

	fmt.Println("append() 来实现 delete操作:")
	//删除s2中的值为8的那个元素
	//append第二个参数是可变参数,所以第二个参数不能直接s2[4:]
	s2 = append(s2[:3], s2[4:]...)
	printSlice2(s2)

	fmt.Println("Poping from front:")
	s2 = s2[1:]
	printSlice2(s2)

	fmt.Println("Poping from back:")
	tail := s2[:len(s2)-1]
	printSlice2(tail)
}

func main() {
	//test1()
	//test2()
	test3()

}
