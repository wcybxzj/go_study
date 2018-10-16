package main

import "fmt"

func test() {
	var a int = 100
	var b bool
	c := 'a'

	fmt.Printf("%+v\n", a)               //100
	fmt.Printf("%#v\n", b)               //false
	fmt.Printf("%T\n", c)                //int32
	fmt.Printf("90%%\n")                 //90%
	fmt.Printf("%t\n", b)                //false
	fmt.Printf("%b\n", 100)              //1100100
	fmt.Printf("%f\n", 199.22)           //199.220000
	fmt.Printf("%q\n", "this is a test") //"this is a test" 注意加了双引号
	fmt.Printf("%x\n", 39839333)         //25fe665 16进制数
	fmt.Printf("%p\n", &a)               //0xc00001e0b8 16进制地址

	str := fmt.Sprintf("a=%d", a)
	fmt.Printf("%q\n", str)
}

/*
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	值的类型的Go语法表示
%%	百分号
*/
func main() {
	test()
}
