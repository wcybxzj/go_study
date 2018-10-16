package main

//定义:
//Go里面有一个关键字iota，这个关键字用来声明enum的时候采用，
//它默认开始值是0，const中每增加一行加1：

//除非被显式设置为其它值或iota，
//每个const分组的第一个常量被默认设置为它的0值，
//第二及后续的常量被默认设置为它前面那个常量的值，
//如果前面那个常量的值是iota，则它也被设置为iota。

import (
	"fmt"
)

// 常量声明省略值时，默认和之前一个值的字面相同。
// 这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
const (
	x = iota
	y = iota
	z = iota
	w
)

// 每遇到一个const关键字，iota就会重置，此时v == 0
const v = iota

//h=0,i=0,j=0 iota在同一行值相同
const (
	h, i, j = iota, iota, iota
)

const (
	a       = iota //0
	b       = "B"
	c       = iota             //2
	d, e, f = iota, iota, iota //都是3
	g       = iota             //4
)

const (
	ybx    = 0 //0
	dabing     //0
)

func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z)
	fmt.Print(ybx, dabing)
}
