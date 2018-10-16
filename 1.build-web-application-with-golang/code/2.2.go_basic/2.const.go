package main

import "fmt"

func test1() {
	const c1 = 123
	const c2 string = "ABC"

	fmt.Print(c1, c2)
	fmt.Print("\n")
}

func main() {
	test1()
}
