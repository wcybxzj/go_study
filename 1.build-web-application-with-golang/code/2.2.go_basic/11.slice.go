package main

import "fmt"

//slice可以简单理解成动态数组
//slice并不是真正意义上的动态数组，而是一个引用类型。
//slice总是指向一个底层array，slice的声明也可以像array一样，只是不需要长度。

func test1() {
	// 和声明array一样，只是少了长度
	//var fslice []int

	//以声明一个slice，并初始化数据
	//slice := []byte{'a', 'b', 'c', 'd'}

}

//slice可以从一个数组或一个已经存在的slice中再次声明。
//slice通过array[i:j]来获取，其中i是数组的开始位置，j是结束位置，
//但不包含array[j]，它的长度是j-i。
func test2() {
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	var a, b []byte
	a = ar[2:5] //2到4
	b = ar[3:5] //3到4
	fmt.Printf("%d %d %d\n", a[0], a[1], a[2])
	fmt.Printf("%d  %d\n", b[0], b[1])
}

//a[:n]
//slice的默认开始位置是0，ar[:n]等价于ar[0:n]

//a[n:]
//slice的第二个序列默认是数组的长度，ar[n:]等价于ar[n:len(ar)]

//a[:]
//从arr直接获取slice，可以ar[:]， 第一个序列是0， 第二个是数组的长度，即等价于ar[0:len(ar)]

func test3() {
	// 声明一个数组
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个slice
	var aSlice, bSlice []byte

	// 演示一些简便操作
	aSlice = array[:3] // 等价于aSlice = array[0:3] aSlice包含元素: a,b,c
	aSlice = array[5:] // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice = array[:]  // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
	aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
	bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
	bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
	bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
	bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g

	fmt.Printf("%c\n", bSlice[0])

}

//slice是引用类型，所以当引用改变其中元素的值时，其它的所有引用都会改变该值，
//例如上面的aSlice和bSlice，如果修改了aSlice中元素的值，那么bSlice相对应的值也会改变。
//
//从概念上面来说slice像一个结构体，这个结构体包含了三个元素：
//一个指针，指向数组中slice指定的开始位置
//长度，即slice的长度
//最大长度，也就是slice开始位置到数组的最后位置的长度
func test4() {
	var arr = [5]byte{11, 22, 33, 44, 55}
	var slice []byte
	var slice2 []byte
	slice = arr[:]
	slice2 = arr[:]
	fmt.Printf("slice[0]:%d\n", slice[0])
	fmt.Printf("slice2[0]:%d\n", slice2[0])
	arr[0] = 123
	fmt.Printf("slice[0]:%d\n", slice[0])
	fmt.Printf("slice2[0]:%d\n", slice2[0])

}

func main() {
	test1()
	test2()
	test3()
	test4()
}
