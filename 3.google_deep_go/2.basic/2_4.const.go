package main

import (
	"fmt"
	"math"
)

const (
	filename string = "abx.txt"
	name            = "ybx"
	aa, bb          = 33, 44
)

func consts() {
	const a, b = 3, 4 //这里可以不指定类型，const就相当于文本替换
	var c int
	c = int(math.Sqrt(a*a + a*a)) //注意这里Sqrt参数不用强转
	fmt.Println(filename, c)
}

func main() {
	consts()
	fmt.Println(aa, bb)
}
