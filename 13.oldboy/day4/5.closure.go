package main

import (
	"fmt"
	"strings"
)

func Adder()func(int)int  {
	var x int
	f := func(d int) int {
		x += d
		return x
	}
	return f
}

func makeSuffix(suffix string) func(string)string {
	f := func(name string) string{
		if strings.HasSuffix(name, suffix)==false {
			return name + suffix
		}
		return name
	}
	return f
}

/*
1
101
1101
*/
func test1()  {
	f := Adder()
	fmt.Println(f(1))
	fmt.Println(f(100))
	fmt.Println(f(1000))
}

func test2(){
	f1 := makeSuffix(".bmp")
	fmt.Println(f1("test"))
	fmt.Println(f1("pic"))

	f2 := makeSuffix(".txt")
	fmt.Println(f2("test"))
	fmt.Println(f2("pic"))
}



func main() {
	//test1()
	test2()
}