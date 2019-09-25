package main

import (
	"container/list"
	"fmt"
	"time"
)

func test1() {
	t := time.Now()

	//初始化
	mylist :=list.New()


	//增加元素
	for i := 0; i<1*1*10; i++ {
		mylist.PushBack(i)
	}

	//打印长度
	fmt.Println("len:",mylist.Len())

	//循环打印
	fmt.Println("for1")
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)

	}
	fmt.Println()

	//删除 bug版本
	//按照设想, 这应该会移除list里所有的元素, 但是, 结果是只移除了第一个.
	//原因是: Remove时候, e.next = nil, 使得for判断中, e != nil不成立了, 所以退出了循环.
	//这时候有两种解决办法:
	for e := mylist.Front(); e != nil; e = e.Next() {
		mylist.Remove(e)
	}

	//循环打印
	fmt.Println("for2")
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()


	//删除 ok版
	var next *list.Element
	for e := mylist.Front(); e != nil; e = next {
		next = e.Next()
		mylist.Remove(e)
	}

	//循环打印
	fmt.Println("for3")
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()


	fmt.Println("list 创建速度: " + time.Now().Sub(t).String())

}

func test2() (*list.List){
	//初始化
	mylist :=list.New()


	//增加元素
	for i := 0; i<1*1*10; i++ {
		mylist.PushBack(i)
	}
	return mylist
}

func main() {

	//test1()

	mylist := test2()
	//循环打印
	fmt.Println("for3")
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value)
	}
	fmt.Println()
}




