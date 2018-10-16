package main

import "fmt"

//包内变量不是全局变量
var (
	aa = 3
	ss = "kkk"
	bb = true
)

//bb:=true//不能怎么写

//1.变量的初始值是0和空,和C不同
func vaiableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s) //q是quote引号的意思
}

func vaiableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func vaiableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func vaiableShorter() {
	//只能函数内怎么使用
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func main() {
	vaiableZeroValue()
	vaiableInitialValue()
	vaiableTypeDeduction()
	vaiableShorter()
	fmt.Println(aa, ss, bb)
}
