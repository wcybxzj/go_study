package main

import (
	"fmt"
)

//1, 1, 2, 3 ,5, 8, 13
//a  b
//   a  b
//      a  b
func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	f := fibonacci()

	fmt.Println(f()) //1
	fmt.Println(f()) //1
	fmt.Println(f()) //2
	fmt.Println(f()) //3
	fmt.Println(f()) //5
	fmt.Println(f()) //8
	fmt.Println(f()) //13
	fmt.Println(f()) //21
}
