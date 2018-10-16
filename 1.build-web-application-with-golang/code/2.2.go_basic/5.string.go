package main

import "fmt"

var hello string
var hello1 string = ""

//变量的声明
func test() {
	no, yes, maybe := "no", "yes", "mamam"
	hello2 := "konichiwa"
	hello1 = "Bonjour"
	fmt.Print(no, yes, maybe)
	fmt.Print("\n")
	fmt.Print(hello1, hello2)
	fmt.Print("\n")
}

//修改string的方法
func test1() {
	var s string = "hello"
	//问题: can not assign to s[0]
	//s[0] = 'c'

	//解决办法1: string to byte 方便转换
	c := []byte(s)
	c[0] = 'c'
	s2 := string(c)
	fmt.Printf("%s\n", s2)

	//解决方法2:
	s = "xxxxx" + s[1:]
	fmt.Printf("%s\n", s)
}

func test2() {
	s := "hello"
	m := "world"
	a := s + m
	fmt.Printf("%s\n", a)
}

func test3() {
	m := `hello
		world`
	fmt.Printf("%s\n", m)
}

func main() {
	//test()
	test1()
	//test2()
	//test3()
}
