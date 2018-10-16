package main

import "fmt"

//正统函数式编程 不能有状态
func adder() func(int) int {
	sum := 0                 //sum自由变量
	return func(v int) int { //v是局部变量
		sum += v
		return sum
	}
}

//比较正统的函数式编程不能使用变量，只能有函数和常量
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func test1() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d = %d\n",
			i, a(i))
	}
}

func test2() {
	var s int
	a := adder2(0)
	for i := 0; i < 10; i++ {
		s, a = a(i)
		fmt.Printf("%d = %d\n",
			i, s)
	}
}

func main() {
	test1()
	//test2()
}
