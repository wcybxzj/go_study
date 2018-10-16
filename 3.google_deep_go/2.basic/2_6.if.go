package main

import (
	"fmt"
	"io/ioutil"
)

//写法1:
func test1() {
	const filename = "abc.txt"
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

//写法2:
func test2() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	//contents和err的定义是在if中所以生存周期,所以外边不能访问
}

func main() {
	//test1()
	test2()
}
