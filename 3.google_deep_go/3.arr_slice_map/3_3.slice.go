package main

import "fmt"

//1.sliece参数不用设置数量,从而避免传参因数量带来的问题 对比3_2.arr.go
//2.slice是可以实现和数组指针一样的效果,像3_2.arr.go的test2()
func updateSlice(s []int) {
	s[0] = 10000
}

/*
go run 3_3.slice.go
s:arr[2:6]: [2 3 4 5]
arr[:6]: [0 1 2 3 4 5]
s1:arr[2:]: [2 3 4 5 6 7]
s2:arr[:]: [0 1 2 3 4 5 6 7]
after updateSlice(s1)
s1:arr[2:]: [10000 3 4 5 6 7]
arr: [0 1 10000 3 4 5 6 7]
after updateSlice(s2)
s2:arr[:]: [10000 1 10000 3 4 5 6 7]
arr: [10000 1 10000 3 4 5 6 7]
reslices s2
s2: [10000 3 4]
*/
func test1() {
	//数组:
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}

	//切片:
	//是对数组的视图
	//切片的取值是半开半闭:
	//意思是 左边的索引是包含的,右变的索引是不包含的
	s := arr[2:6]
	fmt.Println("s:arr[2:6]:", s)    //索引2->5的数据
	fmt.Println("arr[:6]:", arr[:6]) //索引0->5的数据
	s1 := arr[2:]
	fmt.Println("s1:arr[2:]:", s1) //索引2->7的数据
	s2 := arr[:]
	fmt.Println("s2:arr[:]:", s2) //全部数据

	fmt.Println("after updateSlice(s1)")
	updateSlice(s1)
	fmt.Println("s1:arr[2:]:", s1)
	fmt.Println("arr:", arr)

	fmt.Println("after updateSlice(s2)")
	updateSlice(s2)
	fmt.Println("s2:arr[:]:", s2)
	fmt.Println("arr:", arr)

	fmt.Println("reslices s2")
	s2 = s2[:5]
	s2 = s2[2:]
	fmt.Println("s2:", s2)
}

//测试2:slice的扩展-第一个难点
//输出:
//go run 3_3.slice.go
//s1[2 3 4 5] len(s1)=4, cap(s1)=6
//s2[5 6] len(s2)=2, cap(s2)=3
func test2() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6] //2,3,4,5
	s2 := s1[3:5]  //5,6 (没有报错,6都不在s1的范围,s2通过sliece扩展可以取出来这个值)
	// fmt.Println(s1[4]) //报错 索引越界
	fmt.Printf("s1%v len(s1)=%d, cap(s1)=%d\n", s1, len(s1), cap(s1))
	fmt.Printf("s2%v len(s2)=%d, cap(s2)=%d\n", s2, len(s2), cap(s2))
}

func main() {
	//test1()
	test2()
}
