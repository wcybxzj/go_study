package main

import "fmt"

//Go里switch默认每个case最后有break，匹配成功后不会自动向下执行其他case，
//而是跳出整个switch,
func test1() {
	i := 10
	switch i {
	case 1:
		fmt.Println("is 1")
	case 2, 3, 4:
		fmt.Println("is 2,3,4")
	default:
		fmt.Println("is other")
	}
}

//可以使用fallthrough强制执行后面的case代码。
func test2() {
	i := 6
	switch i {
	case 4:
		fmt.Println("<=4")
		fallthrough
	case 5:
		fmt.Println("<=5")
		fallthrough
	case 6:
		fmt.Println("<=6")
		fallthrough
	case 7:
		fmt.Println("<=7")
	default:
		fmt.Println("default")
	}
}

func main() {
	test1()
	test2()
}
