package main

import "fmt"

//由于接口不知道是什么类型,用时需要转成具体类型

//方法1:直接类型转换
func test1()  {
	var t int = 123
	var x interface{}
	x = t
	y :=x.(int)
	fmt.Println(y)
}

//方法2:带判断的类型转换
func test2()  {
	var t int
	var x interface{}
	x = t

	y , ok:=x.(int)
	if ok == false {
		fmt.Println("类型转换失败")
	}
	fmt.Println(y)

	//var g string
	//y , ok=g.(int)
	//if ok == false {
	//	fmt.Println("类型转换失败")
	//}
	//fmt.Println(y)
}


func main() {
	//test1()
	test2()
}
