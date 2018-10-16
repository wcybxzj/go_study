package main

import "fmt"

//传值
func add1(a int) int {
	a = a + 1
	return a
}

//传参
func add2(a *int) int {
	*a = *a + 1
	return *a
}

func test() {
	x := 3
	fmt.Println("x=", x)
	x1 := add1(x)
	fmt.Println("x+1= ", x1)
	fmt.Println("x= ", x)

	fmt.Println("")

	x = 3
	fmt.Println("x=", x)
	x1 = add2(&x)
	fmt.Println("x+1= ", x1)
	fmt.Println("x= ", x)
}

func main() {
	test()
	fmt.Println("")
}
