package main

import "fmt"

func test1() {
	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is ", sum)
}

func test2() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("sum is ", sum)
}

func main() {
	test1()
	test2()
}
