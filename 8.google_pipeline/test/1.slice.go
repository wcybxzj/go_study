package main

import "fmt"

/*
[11 22]
[33 44]
*/
func test1() {
	arr := []int{11, 22, 33, 44}
	m := len(arr) / 2
	fmt.Println(arr[:m])
	fmt.Println(arr[m:])
}

func test2() {
	arr := []int{11, 222}
	num := len(arr) / 2
	fmt.Println(arr[num])
}

func main() {
	test1()
	//test2()
}
