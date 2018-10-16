package main

import (
	"fmt"
)

//现在的理解panic相当

func main() {
	f()
	fmt.Println("return normally from f")

	fmt.Println("========================")

	h()
	fmt.Println("return normally from h")
}

//没错任何处理
func h() {
	fmt.Println("Calling h")
	g()
	fmt.Println("return normally from h")
}

//使用defer+recover处理panic
func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
			fmt.Println("recovered in f", r)
		}
	}()

	fmt.Println("Calling g")
	g()
	fmt.Println("return normally from g")
}

func g() {
	panic("ERROR")
}
