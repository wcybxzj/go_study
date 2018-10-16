package main

import "fmt"

func test() {
	var c complex64 = 5 + 5i
	//output: (5+5i)
	fmt.Printf("Value is: %v", c)
}

func main() {
	test()
}
