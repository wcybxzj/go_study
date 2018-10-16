package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func test1() {
	x := 3
	y := 4
	z := 5
	max_xy := max(x, y)
	max_zx := max(x, z)

	fmt.Printf("%d\n", max_xy)
	fmt.Printf("%d\n", max_zx)
	fmt.Printf("%d\n", max(y, z))
}

func mutl_return1(A, B int) (int, int) {
	return A + B, A * B
}

func mutl_return2(A, B int) (add int, mulitpiled int) {
	add = A + B
	mulitpiled = A * B
	return
}

func test2() {
	x := 3
	y := 4
	v1, v2 := mutl_return1(x, y)
	fmt.Printf("%d\n", v1)
	fmt.Printf("%d\n", v2)

	x = 5
	y = 6
	v1, v2 = mutl_return2(x, y)
	fmt.Printf("%d\n", v1)
	fmt.Printf("%d\n", v2)
}

func test3_inner(arg ...int) {
	for _, n := range arg {
		fmt.Printf("number is %d\n", n)
	}
}

func test3() {
	test3_inner(11, 222)

	my_slice :=[]int{44,55,66}
	test3_inner(my_slice...)
}

func main() {
	test1()
	fmt.Printf("---------------------\n")
	test2()
	fmt.Printf("---------------------\n")
	test3()
}
