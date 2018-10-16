package main

import (
	"fmt"
)

//golang:
//默认每个case 默认带break
//如果不想break,反而要写fallthrough

//测试1:switch 正常情况
func test1(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("err op" + op)
	}
	return result
}

//测试2:switch 为空的情况
func test2(score int) string {
	g := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("wrong score: %d\n", score)) //中断执行
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}
	return g
}

func main() {
	var re int
	re = test1(11, 22, "+")
	fmt.Println(re)

	fmt.Println(
		test2(90),
		test2(80),
		test2(70),
		test2(60),
		test2(50),
		test2(40),
		test2(30),
		test2(20),
		test2(10),
		test2(0),
		//test2(-10),
		//test2(101),
	)

}
