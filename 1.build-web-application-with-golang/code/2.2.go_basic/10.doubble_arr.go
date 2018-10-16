package main

import "fmt"

func test1() {
	a1 := [2][4]int{[4]int{11, 22, 33, 44}, [4]int{55, 66, 77, 88}}
	a2 := [2][4]int{{1, 2, 3, 4}, {5, 6, 6, 7}}
	fmt.Printf("%d\n", a1[0][0])
	fmt.Printf("%d\n", a2[0][0])
}

func main() {
	test1()
}
