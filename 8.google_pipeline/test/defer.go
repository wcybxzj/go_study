package main

import "fmt"

func t1() {
	fmt.Println("111111111")
}

func t2() {
	fmt.Println("222222222")
}

/*
输出:
222222222222
11111111111
*/
func main() {
	defer t1()
	defer t2()
}
