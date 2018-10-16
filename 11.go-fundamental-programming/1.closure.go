package main

import "fmt"

/*
4
3
2
1
*/
func test1() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	defer fmt.Println("4")
}

/*
closure i =  4
closure i =  4
closure i =  4
closure i =  4
defer_closure i =  4
defer_closure i =  4
defer_closure i =  4
defer_closure i =  4
*/
func test2() {
	var fs = [4]func(){}

	for i := 0; i < 4; i++ {
		defer func() { fmt.Println("defer_closure i = ", i) }()
		fs[i] = func() { fmt.Println("closure i = ", i) }
	}

	for _, f := range fs {
		f()
	}
}

func main() {
	//test1()
	test2()
}
