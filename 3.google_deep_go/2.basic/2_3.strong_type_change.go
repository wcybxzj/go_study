package main

import (
	"fmt"
	"math"
)

//golang中只有强类型转换,没有隐式类型转换
func triangle() {
	a, b := 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func main() {
	triangle()
}
