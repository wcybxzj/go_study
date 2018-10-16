package main

import "fmt"

var isActive bool                   //全局变量
var enabled, disabled = true, false //忽略类型的声明

func test() {
	var available bool //一般声明
	available = true   //赋值操作
	vaild := false     //简短声明

	fmt.Print(available)
	fmt.Print(vaild)
}

func main() {
	test()
}
